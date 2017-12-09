(function() {
    var js = $('[src="assets/js/core/ReportingTool.js"]');
    js.before("<script src='assets/js/core/jquery.history.js' type='text/javascript'></script>");
    js.before("<script src='assets/plugins/bootstrap-datepicker/bootstrap-datepicker.min.js' type='text/javascript'></script>");
    js.before("<script src='assets/plugins/bootstrap-datepicker/bootstrap-datepicker.zh-CN.min.js' type='text/javascript'></script>");
    // js.before("<script src='/ims_system/js/TableInModal.js' type='text/javascript'></script>");
    var css = $('[href="assets/css/reportingtool/ReportingTool.css"]');
    css.before("<link href='assets/plugins/bootstrap-datepicker/bootstrap-datepicker.css' rel='stylesheet' type='text/css' />");
    // css.before("<link href='/ims_system/css/TableInModal.css' rel='stylesheet' type='text/css' />");
    if ($.support.opacity) {
        css.after("<link href='assets/css/reportingtool/CheckBox.css' rel='stylesheet' type='text/css' />");
    }
})();

(function($) {
    var serverURL = "user/report/";

    var cachedRows = {};
    var nullRows = {};

    $.fn.rt = function(options) {

        var _this = this,
            globalVars = {
                tableID: "",
                pageTitle: "",
                initUrl: "",
                url: "",
                queryObj: { table: "", page: 1, rows: 15, colpage: 1 }
            };

        var settings = $.extend({
            asyncLoad: false,
            asyncRefresh: $.support.opacity,
            complete: function() {},
            configFile: "",
            hasCheckbox: false,
            hasPager: true,
            navBar: false,
            nodeLinker: true,
            query: "",
            rowList: [15, 30, 60, 100],
            saveState: true,
            searchBar: true,
            searchBarOnShow: function() {},
            searchBarOnHide: function() {},
            searchComplete: function() {},
            striped: false,
            style: "table",
            table: "",
        }, options);

        this.addClass("rt-content");

        this.html("<ol class=\"rt-nav breadcrumb\"></ol>\
                      <div class=\"rt-condition\"></div>\
                      <div class=\"rt-selector\"></div>\
                      <div class=\"rt-search rt-" + settings.style + "\" style=\"display:none\"></div>\
                      <div class=\"rt-body rt-" + settings.style + "\"></div>");

        var tableSearcher = "<div class=\"rt-search-cdts\"></div>\
                                    <div class=\"rt-search-btns\">\
                                        <span class=\"rt-search-showadv\">\
                                            <span class=\"rt-search-showadv-txt\">更多</span><span class=\"glyphicon glyphicon-chevron-down rt-glyphicon-color\"></span>\
                                        </span>\
                                        <span class=\"rt-search-go\">\
                                            查询<span class=\"glyphicon glyphicon-search rt-glyphicon-color\"></span>\
                                        </span>\
                                    </div>";

        var treeSearcher = "<div class=\"rt-treeSearcher-div\">\
                                       <input type=\"text\" class=\"rt-treeSearcher-txt\"/>\
                                       <span class=\"glyphicon glyphicon-search rt-treeSearcher-btn rt-glyphicon-color\"></span>\
                                   </div>";

        var css = $('head>[href*="/assets/css/reportingtool/ReportingTool.css"]');
        if (settings.nodeLinker) {
            css.after("<link href='/assets/css/reportingtool/NodeLinker.css' rel='stylesheet' type='text/css' />");
        }

        var getSettingsQueryVar = function(variable) {
            var query = settings.query.substring(1);
            var vars = query.split("&");
            for (var i = 0; i < vars.length; i++) {
                var pair = vars[i].split("=");
                if (pair[0] == variable) {
                    return decodeURI(pair[1]);
                }
            }
            return (false);
        };
        var getQueryVariable = function(variable) {
            var query = window.location.search.substring(1);
            var vars = query.split("&");
            for (var i = 0; i < vars.length; i++) {
                var pair = vars[i].split("=");
                if (pair[0] == variable) {
                    return decodeURI(pair[1]);
                }
            }
            return (false);
        };

        globalVars.pageTitle = $("title").html();

        var table = "";
        if (table = getSettingsQueryVar("table")) {
            globalVars.queryObj.table = encodeURI(table);
        } else if (table = getQueryVariable("table")) {
            globalVars.queryObj.table = encodeURI(table);
        } else {
            globalVars.queryObj.table = encodeURI(settings.table);
        }
        globalVars.queryObj.rows = settings.rowList[0];

        var getQuery = function() {
            if (settings.query) {
                return settings.query;
            } else {
                return location.search;
            }
        };
        var load = function(elem) {
            var queryStr = getQuery();
            if (queryStr.indexOf("page") === -1) {
                if (queryStr === "") {
                    queryStr += "?";
                } else {
                    queryStr += "&";
                }
                queryStr += "page=" + globalVars.queryObj.page + "&rows=" + globalVars.queryObj.rows + "&colpage=" + globalVars.queryObj.colpage;
            }
            if (queryStr.indexOf("table") === -1) {
                queryStr += "&table=" + globalVars.queryObj.table;
            }
            if (settings.query) {
                settings.query = queryStr;
                postData(elem);
            } else if (settings.asyncRefresh) {
                History.replaceState(null, globalVars.pageTitle, queryStr);
                postData(elem);
            } else {
                var hr = getQueryVariable("hr");
                if (!hr) {
                    queryStr += "&hr=true";
                    location.href = location.pathname + queryStr;
                } else {
                    postData(elem);
                }
            }
        };
        var format = function() {
            var condition = $(".rt-condition");
            if (!condition.find("div").length) {
                condition.css("display", "none");
            } else {
                condition.css("display", "block");
            }
            var selector = $(".rt-selector");
            if (!selector.find("div").length) {
                selector.css("display", "none");
            } else {
                selector.css("display", "block");
            }
        };
        var postData = function(elem) {
            var postOpts = {
                async: settings.asyncLoad,
                method: "POST",
                url: serverURL + "GetTable" + getQuery(),
                data: {
                    configFile: settings.configFile,
                    hasCheckbox: settings.hasCheckbox,
                    style: settings.style,
                    rowList: settings.rowList.toString()
                },
                success: function(data) {
                    if (data.status === "fail") {
                        alert(data.msg);
                        location.href = location.pathname;
                    }
                    var jsonObject = data.res;
                    if (settings.searchBar === true) {
                        var rtSearch = _this.find(".rt-search");
                        if (settings.style != "tree") {
                            rtSearch.html(tableSearcher);
                            rtSearch.find(".rt-search-cdts").html(jsonObject.search);
                            if (jsonObject.search) {
                                rtSearch.css("display", "block");
                            } else {
                                rtSearch.css("display", "none");
                            }
                        }
                    }
                    if (jsonObject.condition) {
                        _this.find(".rt-condition").html(jsonObject.condition);
                    } else {
                        _this.find(".rt-condition").css("display", "none");
                    }
                    if (jsonObject.selector) {
                        _this.find(".rt-selector").html(jsonObject.selector);
                    } else {
                        _this.find(".rt-selector").css("display", "none");
                    }
                    _this.find(".rt-body").html(jsonObject.body);
                    if (settings.striped) {
                        _this.find("table").addClass("table-striped");
                    }
                    if (!settings.hasPager) {
                        _this.find(".rt-pager-container").css("display", "none");
                    }
                    var nav = _this.find(".rt-nav");
                    if (settings.searchBar === true && settings.style === "tree") {
                        nav.append(treeSearcher);
                    }
                    if (settings.navBar === true) {
                        if (!jsonObject.exception) {
                            nav.append("<li><a class='rt-nav-a' href='#' data-url='" + getQuery() + "'>" + decodeURI($(elem).attr("data-navname")) + "</a></li>");
                        }
                        if (elem === undefined) {
                            nav.find("a").eq(0).attr("data-url", getQuery()).html(decodeURI(globalVars.queryObj.table));
                        }
                    } else {
                        if (settings.searchBar === true && settings.style === "tree") {
                            _this.find(".rt-nav>li").css("display", "none");
                        } else {
                            _this.find(".rt-nav").css("display", "none");
                        }
                    }
                    _this.off();
                    _this.on("click", ".rt-search-showadv", showAdvSearch).on("click", ".rt-search-go", search).on("click", ".rt-sort", sort);
                    _this.on("click", ".rt-selector-selectmore", selectmore).on("click", ".rt-selector-multiselect", multiselect);
                    _this.on("click", ".rt-multiselect-cancel", cancelMultiselect).on("click", ".rt-multiselect-ok", submitMultiselect);
                    _this.on("click", ".rt-selector-list-text", selectCondition).on("click", ".rt-condition-remove", removeCondition);
                    _this.on("click", ".rt-pager-firstPage", firstPage).on("click", ".rt-pager-lastPage", lastPage);
                    _this.on("click", ".rt-pager-prevPage", prevPage).on("click", ".rt-pager-nextPage", nextPage);
                    _this.on("click", ".rt-colPager-prev", prevCols).on("click", ".rt-colPager-next", nextCols);
                    _this.on("keypress", ".rt-pager-page", page_Keypress).on("change", ".rt-pager-rowList", rowList_Change);
                    _this.on("mouseover", "tbody tr", trOnMouseover).on("mouseout", "tbody tr", trOnMouseout);
                    _this.on("click", "tbody tr", trOnClick).on("click", ".rt-th-checkbox>.rt-checkboxWrapper", toggleAll);
                    _this.on("click", ".rt-pager-export", exportExcel);
                    _this.on("keyup", ".rt-celltext", updateCellValue).on("change", ".rt-cellselect", updateCellValue);
                    _this.on("click", "[data-table]", getTable).on("click", ".rt-nav-a", getTableByNav);
                    _this.on("click", ".rt-node>.glyphicon-triangle-right", expandNode); //.on("dblclick", "[data-childtree]", expandNode);
                    _this.on("click", ".rt-treeSearcher-btn", searchTree).on("click", ".rt-search-result .rt-node-cols", locateNode);
                    _this.on("click", ".rt-node-cols", nodeOnClick);
                    _this.on("click", ".rt-node>.rt-checkboxWrapper", checkNode);
                    _this.on("click", ".rt-create", createOne).on("click", ".rt-view", viewThis).on("click", ".rt-edit", editThis).on("click", ".rt-delete", deleteThis)
                        //_this.on("keyup", ".rt-search-txt", startSearching).on("change", ".rt-search-txt.date", startSearching);
                    $(".rt-search-txt.date").datepicker({
                        format: "yy/mm/dd",
                        weekStart: 1,
                        language: "zh-CN",
                        orientation: "bottom left",
                        keyboardNavigation: false,
                        autoclose: true,
                        todayHighlight: true
                    });
                    var queryArray = getQuery().substring(1).split('&');
                    for (var i = 0; i < queryArray.length; i++) {
                        var keyValuePair = queryArray[i].split('=');
                        var key = keyValuePair[0];
                        var sign = /%\d[A-Za-z]%\d[A-Za-z]/.exec(key);
                        key = sign && key.substring(0, key.length - 6) + sign.toString().toLocaleLowerCase();
                        globalVars.queryObj[key || keyValuePair[0]] = keyValuePair[1];
                    }
                    $(".rt-condition>div").each(function() {
                        var key = $(this).attr("data-value");
                        $(".rt-search-cdts").find("[name=\"" + key + "\"]").closest("div").css("display", "none");
                    });
                    globalVars.tableID = $(_this).attr("id") || Math.floor(Math.random() * (10000000 - 0)) + 0;
                    cachedRows[globalVars.tableID] = {};
                    nullRows[globalVars.tableID] = jsonObject.row;
                    if (!jsonObject.exception) {
                        settings.complete();
                    }
                },
                error: function() {
                    alert("您未搭建服务器哦！");
                    return false;
                }
            };
            $.ajax(postOpts);
        };
        var getTable = function() {
            var elem = this;
            globalVars.queryObj.table = encodeURI($(this).attr("data-table"));
            if (settings.query) {
                settings.query = "?" + $(this).attr("data-passedcol");
            } else {
                History.replaceState(null, globalVars.pageTitle, "?" + $(this).attr("data-passedcol"));
            }
            load(elem);
        };
        var getTableByNav = function() {
            var elem = this;
            var url = $(this).attr("data-url");
            if (settings.query) {
                settings.query = url;
            } else {
                History.replaceState(null, globalVars.pageTitle, url);
            }
            $.post(serverURL + "GetTable" + getQuery(), {
                configFile: settings.configFile,
                hasCheckbox: settings.hasCheckbox,
                style: settings.style,
                rowList: settings.rowList.toString()
            }, function(data) {
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    location.href = location.pathname;
                }
                var jsonObject = data.res;
                if (settings.searchBar === true) {
                    if (jsonObject.search) {
                        if (settings.style === "table") {
                            var rtSearch = _this.find(".rt-search");
                            rtSearch.html(tableSearcher);
                            rtSearch.find(".rt-search-cdts").html(jsonObject.search);
                            rtSearch.css("display", "block");
                        }
                    }
                }
                _this.find(".rt-condition").html(jsonObject.condition);
                _this.find(".rt-selector").html(jsonObject.selector);
                _this.find(".rt-body").html(jsonObject.body);
                if (settings.striped) {
                    _this.find("table").addClass("table-striped");
                }
                if (!settings.hasPager) {
                    _this.find(".rt-pager-container").css("display", "none");
                }
                $(elem).parent().nextAll().remove();
                if (!jsonObject.exception) {
                    settings.complete();
                }
                format();
            });
        };
        var buildQueryStr = function() {
            var queryStr = "",
                value = "";
            for (var q in globalVars.queryObj) {
                value = globalVars.queryObj[q];
                if (value.toString().indexOf("%") === -1) {
                    value = encodeURI(value);
                }
                queryStr += "&" + q + "=" + value;
            }
            return "?" + queryStr.substring(1);
        };
        var refresh = function() {
            if (settings.query) {
                settings.query = buildQueryStr();
            } else {
                History.replaceState(null, globalVars.pageTitle, buildQueryStr());
            }
            $.post(serverURL + "GetTable" + getQuery(), {
                configFile: settings.configFile,
                hasCheckbox: settings.hasCheckbox,
                style: settings.style,
                rowList: settings.rowList.toString()
            }, function(data) {
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    location.href = location.pathname;
                }
                var jsonObject = data.res;
                _this.find(".rt-body").html(jsonObject.body);
                _this.find(".rt-condition").html(jsonObject.condition);
                _this.find(".rt-selector").html(jsonObject.selector);
                if (settings.striped) {
                    _this.find("table").addClass("table-striped");
                }
                if (!settings.hasPager) {
                    _this.find(".rt-pager-container").css("display", "none");
                }
                if (!jsonObject.exception) {
                    settings.complete();
                }
                $("td .rt-checkboxWrapper").each(function() {
                    var checkbox = $(this).find(".rt-checkbox");
                    if (cachedRows[globalVars.tableID][checkbox.val()]) {
                        checkbox[0].checked = true;
                        $(this).addClass("checked");
                        $(this).closest("tr").addClass("rt-tr-selected");
                    }
                });
                var hasCheckbox = _this.find("td .rt-checkbox").length;
                if (hasCheckbox && hasCheckbox === _this.find("td .rt-checkbox:checked").length) {
                    var thCheckbox = _this.find(".rt-th-checkbox");
                    thCheckbox.find(".rt-checkboxWrapper").addClass("checked");
                    thCheckbox.find(".rt-checkbox")[0].checked = true;
                }
                format();
            });
        };
        var expandNode = function() {
            var parentNode = $(this).parent();
            var arrow = parentNode.children(".glyphicon-triangle-right");
            if (arrow.hasClass("down")) {
                var childtree = parentNode.children(".rt-childtree");
                childtree.slideUp("fast");
                arrow.removeClass("down");
            } else {
                parentNode.children(".rt-childtree").remove();
                parentNode.append("<div class='rt-childtree' style='display:none'></div>");
                var childtree = parentNode.children(".rt-childtree");
                var slideDown = function() {
                    if (childtree.siblings(".rt-checkboxWrapper").hasClass("checked")) {
                        childtree.find(".rt-checkboxWrapper").each(function() {
                            $(this).addClass("checked");
                            $(this).find(".rt-checkbox")[0].checked = true;
                        });
                    }
                    childtree.slideDown("fast");
                };
                var trees = JSON.parse(parentNode.children(".rt-node-cols").attr("data-childtree"));
                var treesLength = Object.keys(trees).length,
                    counter = 0;
                var successfunc = function(data) {
                    // var jsonObject = JSON.parse(data);
                    if (data.status === "fail") {
                        alert(data.msg);
                        location.href = location.pathname;
                    }
                    var jsonObject = data.res;
                    childtree.append(jsonObject.body);
                    counter++;
                    if (counter === treesLength) {
                        slideDown();
                        arrow.addClass("down");
                    }
                };
                for (var t in trees) {
                    $.ajax({
                        async: false,
                        url: serverURL + "GetTable" + "?table=" + encodeURI(t) + "&" + trees[t],
                        method: "POST",
                        data: {
                            configFile: settings.configFile,
                            hasCheckbox: settings.hasCheckbox,
                            style: settings.style
                        },
                        success: successfunc
                    });
                }
            }
        };
        var searchTree = function() {
            var cd = $(".rt-treeSearcher-txt").val();
            if (!cd) {
                return;
            }
            $.post(serverURL + "SearchTree" + "?table=" + globalVars.queryObj.table, {
                configFile: settings.configFile,
                hasCheckbox: settings.hasCheckbox,
                style: settings.style,
                condition: cd
            }, function(data) {
                if (data.status === "fail") {
                    alert(data.msg);
                    location.href = location.pathname;
                }
                var jsonObject = data.res;
                // var jsonObject = JSON.parse(data);
                _this.find(".rt-body").html(jsonObject.body);
            });
        };
        var locateNode = function() {
            var table = encodeURI($(this).attr("data-tableid"));
            $.post(serverURL + "LocateNode" + "?table=" + table, {
                configFile: settings.configFile,
                hasCheckbox: settings.hasCheckbox,
                style: settings.style,
                condition: $(this).attr("data-parentnode")
            }, function(data) {
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    location.href = location.pathname;
                }
                var jsonObject = data.res;
                for (var i = 0; i < jsonObject.length; i++) {
                    var obj = jsonObject[i],
                        elem = null;
                    if (obj.parent != "ROOTNODE") {
                        var selector = ".rt-node-cols[data-childtree='{" + obj.parent.replace(/"/g, "\"") + "}']";
                        var parent = _this.find(selector).parent();
                        parent.append("<div class=\"rt-childtree\"></div>");
                        parent.children(".rt-childtree").html(obj.elems);
                        parent.children(".glyphicon-triangle-right").addClass("down");
                    } else {
                        _this.find(".rt-body").html(obj.elems);
                    }
                }
                var a = jsonObject;
            });
        };
        var search = function() {
            $(".rt-search-txt").each(function() {
                var value = $.trim($(this).val());
                var name = $(this).attr("name"),
                    sign = $(this).attr("data-sign");
                if (value) {
                    globalVars.queryObj[(name + sign)] = value;
                } else {
                    delete globalVars.queryObj[(name + sign)];
                }
            });
            if (settings.query || settings.asyncRefresh) {
                refresh();
            } else {
                location.href = location.pathname + buildQueryStr();
            }
            settings.searchComplete();
            return false;
        };
        var createOne = function() {
            $.post(serverURL + "GetPageCRUD" + "?table=" + globalVars.queryObj.table, {
                cmd: "add",
                configFile: settings.configFile
            }, function(data) {
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    location.href = location.pathname;
                }
            });
        };
        var viewThis = function() {
            $.post(serverURL + "GetPageCRUD" + "?table=" + globalVars.queryObj.table, {
                cmd: "view",
                configFile: settings.configFile
            }, function(data) {
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    location.href = location.pathname;
                }
            });
        };
        var editThis = function() {
            $.post(serverURL + "GetPageCRUD" + "?table=" + globalVars.queryObj.table, {
                cmd: "edit",
                configFile: settings.configFile
            }, function(data) {
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    location.href = location.pathname;
                }
            });
        };
        var deleteThis = function() {
            $.post(serverURL + "GetPageCRUD" + "?table=" + globalVars.queryObj.table, {
                cmd: "delete",
                configFile: settings.configFile
            }, function(data) {
                // var jsonObject = JSON.parse(data);
                if (data.status === "fail") {
                    alert(data.msg);
                    location.href = location.pathname;
                }
            });
        };
        var startSearching = function() {
            setTimeout(search, 1);
        };
        var clean = function() {
            $(".rt-search-txt").each(function() {
                $(this).val("");
            });
            search();
            return false;
        };
        var showAdvSearch = function() {
            if ($(this).hasClass("adv-shown")) {
                $(".rt-search-adv").css("display", "none");
                $(this).find(".rt-search-showadv-txt").html("更多");
                $(this).find(".glyphicon-chevron-down").removeClass("upsidwn");
                $(this).removeClass("adv-shown");
            } else {
                var conditionBlock = $(".rt-condition");
                $(".rt-search-adv").each(function() {
                    var attrValue = $(this).find(".rt-search-txt").attr("name");
                    if (!conditionBlock.find("[data-value=\"" + attrValue + "\"]").length) {
                        $(this).css("display", "inline-block");
                    }
                });
                $(this).find(".rt-search-showadv-txt").html("收起");
                $(this).find(".glyphicon-chevron-down").addClass("upsidwn");
                $(this).addClass("adv-shown");
            }
        };
        var sort = function() {
            var sortIcon = $(this).find(".glyphicon");
            if (sortIcon.length && sortIcon.hasClass("glyphicon-arrow-up")) {
                globalVars.queryObj.sort = $(this).attr("name") + "%20DESC";
            } else {
                globalVars.queryObj.sort = $(this).attr("name") + "%20ASC";
            }
            if (settings.query || settings.asyncRefresh) {
                refresh();
            } else {
                location.href = location.pathname + buildQueryStr();
            }
        };
        var selectmore = function() {
            var list = $(this).parent().siblings(".rt-selector-value").find(".rt-selector-list");
            if ($(this).hasClass("selectmore-all")) {
                list.css("height", "30px");
                $(this).find(".rt-selectmore-txt").html("更多");
                $(this).find(".glyphicon-chevron-down").removeClass("upsidwn");
                $(this).removeClass("selectmore-all");
            } else {
                list.css("height", "auto");
                $(this).find(".rt-selectmore-txt").html("收起");
                $(this).find(".glyphicon-chevron-down").addClass("upsidwn");
                $(this).addClass("selectmore-all");
            }
        };
        var cancelMultiselect = function() {
            $(".rt-multiselect-btns").css("display", "none");
            $(".rt-selector-list-text>.glyphicon").css("display", "none");
            $(".rt-multiselect-btns>button").css("display", "none");
            $(".rt-multiselect-checked").removeClass("rt-multiselect-checked");
            $(".rt-selector-list-text>.glyphicon-check").removeClass("glyphicon-check").addClass("glyphicon-unchecked");
            $(this).parent().siblings(".rt-selector-btns").find(".rt-selector-multiselect").removeClass("rt-multiselect-true");
        };
        var multiselect = function() {
            cancelMultiselect();
            var elem = $(this).parent().siblings(".rt-selector-value");
            var list = elem.find(".rt-selector-list");
            var sm = $(this).siblings(".rt-selector-selectmore");
            var btns = elem.siblings(".rt-multiselect-btns");
            list.css("height", "auto");
            sm.find(".rt-selectmore-txt").html("收起");
            sm.find(".glyphicon-chevron-down").addClass("upsidwn");
            sm.addClass("selectmore-all");
            list.find(".rt-selector-list-text>.glyphicon").css("display", "inline-block");
            btns.css("display", "block");
            btns.find(".rt-multiselect-cancel").css("display", "inline-block");
            $(this).addClass("rt-multiselect-true");
        };
        var submitCondition = function(key) {
            var conditions = [];
            $(".rt-selector-key[data-value=\"" + key + "\"]").siblings(".rt-selector-value").find(".rt-multiselect-checked").each(function() {
                conditions.push($(this).attr("data-value"));
            });
            var conditionStr = conditions.join("|");
            globalVars.queryObj[key] = conditionStr;
            if (settings.query || settings.asyncRefresh) {
                refresh();
            } else {
                location.href = location.pathname + buildQueryStr();
            }
            $(".rt-search-cdts").find("[name=\"" + key + "\"]").closest("div").css("display", "none");
        };
        var submitMultiselect = function() {
            var key = $(this).parent().siblings(".rt-selector-key").attr("data-value");
            submitCondition(key);
        };
        var selectCondition = function() {
            var elem = $(this).closest(".rt-selector-value");
            var multiselect = elem.siblings(".rt-selector-btns").find(".rt-selector-multiselect");
            if (!multiselect.hasClass("rt-multiselect-true")) {
                var key = $(this).closest(".rt-selector-value").siblings(".rt-selector-key").attr("data-value");
                $(this).parent().addClass("rt-multiselect-checked");
                $(".rt-search-cdts").find("[name=\"" + key + "\"]").closest("div").css("display", "none");
                submitCondition(key);
            } else {
                var checkbox = $(this).find(".glyphicon");
                if (checkbox.hasClass("glyphicon-unchecked")) {
                    checkbox.removeClass("glyphicon-unchecked");
                    checkbox.addClass("glyphicon-check");
                    $(this).parent().addClass("rt-multiselect-checked");
                } else {
                    checkbox.removeClass("glyphicon-check");
                    checkbox.addClass("glyphicon-unchecked");
                    $(this).parent().removeClass("rt-multiselect-checked");
                }
                if (elem.find(".rt-multiselect-checked").length) {
                    elem.siblings(".rt-multiselect-btns").find(".rt-multiselect-ok").css("display", "inline-block");
                } else {
                    elem.siblings(".rt-multiselect-btns").find(".rt-multiselect-ok").css("display", "none");
                }
            }
        };
        var removeCondition = function() {
            var key = $(this).parent().attr("data-value");
            var txt = $(".rt-search-cdts").find("[name=\"" + key + "\"]");
            txt.val("");
            var searchDiv = txt.parent("div");
            if (!searchDiv.hasClass("rt-search-adv")) {
                searchDiv.css("display", "inline-block");
            } else if (searchDiv.parent("div").siblings(".rt-search-btns").find(".rt-search-showadv").hasClass("adv-shown")) {
                searchDiv.css("display", "inline-block");
            }
            delete globalVars.queryObj[key];
            if (settings.query || settings.asyncRefresh) {
                refresh();
            } else {
                location.href = location.pathname + buildQueryStr();
            }
        };
        var updateCellValue = function() {
            var value = $(this).val();
            $(this).attr("value", value);
            $(this).parent().attr("data-value", value);
        };
        var firstPage = function() {
            globalVars.queryObj.page = 1;
            if (settings.query || settings.asyncRefresh) {
                refresh();
            } else {
                location.href = location.pathname + buildQueryStr();
            }
        };
        var prevPage = function() {
            var pageNumber = Number(_this.find(".rt-pager-page").val());
            if (pageNumber > 1) {
                globalVars.queryObj.page = pageNumber - 1;
                if (settings.query || settings.asyncRefresh) {
                    refresh();
                } else {
                    location.href = location.pathname + buildQueryStr();
                }
            }
        };
        var nextPage = function() {
            var pageNumber = Number(_this.find(".rt-pager-page").val());
            var totalPage = Number(_this.find(".rt-pager-totalPages").html());
            if (pageNumber < totalPage) {
                globalVars.queryObj.page = pageNumber + 1;
                if (settings.query || settings.asyncRefresh) {
                    refresh();
                } else {
                    location.href = location.pathname + buildQueryStr();
                }
            }
        };
        var lastPage = function() {
            globalVars.queryObj.page = _this.find(".rt-pager-totalPages").html();
            if (settings.query || settings.asyncRefresh) {
                refresh();
            } else {
                location.href = location.pathname + buildQueryStr();
            }
        };
        var page_Keypress = function() {
            if (event.keyCode === 13) {
                var pageNumber = parseInt(_this.find(".rt-pager-page").val());
                var totalPage = Number(_this.find(".rt-pager-totalPages").html());
                if (isNaN(pageNumber) || pageNumber < 1) {
                    pageNumber = 1;
                } else if (pageNumber > totalPage) {
                    pageNumber = totalPage;
                }
                globalVars.queryObj.page = pageNumber;
                if (settings.query || settings.asyncRefresh) {
                    refresh();
                } else {
                    location.href = location.pathname + buildQueryStr();
                }
            }
        };
        var rowList_Change = function() {
            var pageNumber = parseInt(_this.find(".rt-pager-page").val());
            var rowsPerPage = parseInt($(this).val());
            var totalRecords = parseInt(_this.find(".rt-pager-totalRecords").html());
            if (pageNumber * rowsPerPage > totalRecords) {
                var x = ~~(totalRecords / rowsPerPage);
                var y = totalRecords % rowsPerPage == 0 ? 0 : 1;
                globalVars.queryObj.page = x + y;
            }
            globalVars.queryObj.rows = rowsPerPage;
            if (settings.query || settings.asyncRefresh) {
                refresh();
            } else {
                location.href = location.pathname + buildQueryStr();
            }
        };
        var prevCols = function() {
            var pageNumber = Number(_this.find(".rt-colPager-page").val());
            if (pageNumber > 1) {
                globalVars.queryObj.colpage = pageNumber - 1;
                if (settings.query || settings.asyncRefresh) {
                    refresh();
                } else {
                    location.href = location.pathname + buildQueryStr();
                }
            }
        };
        var nextCols = function() {
            var pageNumber = Number(_this.find(".rt-colPager-page").val());
            var totalPage = Number(_this.find(".rt-colPager-totalColPages").val());
            if (pageNumber < totalPage) {
                globalVars.queryObj.colpage = pageNumber + 1;
                if (settings.query || settings.asyncRefresh) {
                    refresh();
                } else {
                    location.href = location.pathname + buildQueryStr();
                }
            }
        };
        var trOnMouseover = function() {
            $(this).addClass("rt-tr-onhover");
        };
        var trOnMouseout = function() {
            $(this).removeClass("rt-tr-onhover");
        };
        var trOnClick = function() {
            var elem = $(this).find(".rt-td-checkbox");
            if (elem.length) {
                if ($(this).hasClass("rt-tr-selected")) {
                    $(this).removeClass("rt-tr-selected");
                    elem.find(".rt-checkboxWrapper").removeClass("checked");
                    var checkbox = elem.find(".rt-checkbox");
                    if (checkbox.length) {
                        checkbox[0].checked = false;
                        var rowID = checkbox.val();
                        delete cachedRows[globalVars.tableID][rowID];
                        var thCheckbox = _this.find(".rt-th-checkbox");
                        thCheckbox.find(".rt-checkboxWrapper").removeClass("checked");
                        thCheckbox.find(".rt-checkbox")[0].checked = false;
                    }
                } else {
                    $(this).addClass("rt-tr-selected");
                    elem.find(".rt-checkboxWrapper").addClass("checked");
                    var checkbox = elem.find(".rt-checkbox");
                    if (checkbox.length) {
                        checkbox[0].checked = true;
                        var rowid = checkbox.val();
                        if (!cachedRows[globalVars.tableID][rowid]) {
                            var rowObj = {};
                            var cells = $(this).find("td");
                            cells.each(function() {
                                rowObj[$(this).attr("name")] = $(this).attr("data-value");
                            });
                            cachedRows[globalVars.tableID][rowid] = rowObj;
                        }
                        if (_this.find("td .rt-checkbox").length === _this.find("td .rt-checkbox:checked").length) {
                            var thCheckbox = _this.find(".rt-th-checkbox");
                            thCheckbox.find(".rt-checkboxWrapper").addClass("checked");
                            thCheckbox.find(".rt-checkbox")[0].checked = true;
                        }
                    }
                }
            } else {
                var cells = $(this).find("td");
                if ($(this).hasClass("rt-tr-selected")) {
                    $(this).removeClass("rt-tr-selected");
                    var rowid = cells.eq(0).attr("data-value");
                    delete cachedRows[globalVars.tableID][rowid];
                } else {
                    cachedRows[globalVars.tableID] = {};
                    var selected = _this.find(".rt-tr-selected");
                    if (selected.length) {
                        selected.removeClass("rt-tr-selected");
                    }
                    $(this).addClass("rt-tr-selected");
                    var rowid = cells.eq(0).attr("data-value");
                    if (!cachedRows[globalVars.tableID][rowid]) {
                        var rowObj = {};
                        cells.each(function() {
                            rowObj[$(this).attr("name")] = $(this).attr("data-value");
                        });
                        cachedRows[globalVars.tableID][rowid] = rowObj;
                    }
                }
            }
        };
        var toggleAll = function() {
            var allCheckbox = _this.find("td .rt-checkboxWrapper");
            var hasChecked = $(this).hasClass("checked");
            if (hasChecked) {
                $(this).removeClass("checked");
                $(this).find(".rt-checkbox")[0].checked = false;
                _this.find(".rt-tr-selected").removeClass("rt-tr-selected");
                allCheckbox.removeClass("checked");
                allCheckbox.each(function() {
                    var checkbox = $(this).find(".rt-checkbox");
                    checkbox[0].checked = false;
                    var rowID = checkbox.val();
                    delete cachedRows[globalVars.tableID][rowID];
                });
            } else {
                $(this).addClass("checked");
                $(this).find(".rt-checkbox")[0].checked = true;
                _this.find("tbody tr").addClass("rt-tr-selected");
                allCheckbox.addClass("checked");
                allCheckbox.each(function() {
                    var checkbox = $(this).find(".rt-checkbox");
                    checkbox[0].checked = true;
                    var rowid = checkbox.val();
                    if (!cachedRows[globalVars.tableID][rowid]) {
                        var rowObj = {};
                        var cells = $(this).closest("tr").find("td");
                        cells.each(function() {
                            rowObj[$(this).attr("name")] = $(this).attr("data-value");
                        });
                        cachedRows[globalVars.tableID][rowid] = rowObj;
                    }
                });
            }
        };
        var checkParent = function(checkbox) {
            var tree = checkbox.closest(".rt-childtree");
            var node = tree.children(".rt-node");
            var allSiblings = node.children(".rt-checkboxWrapper");
            var checkedSiblings = node.children(".rt-checkboxWrapper.checked");
            if (allSiblings.length && allSiblings.length === checkedSiblings.length) {
                var parentCheckbox = tree.siblings(".rt-checkboxWrapper");
                parentCheckbox.addClass("checked");
                parentCheckbox.find(".rt-checkbox")[0].checked = true;
                checkParent(parentCheckbox);
            }
        };
        var checkNode = function() {
            var childCheckbox = $(this).parent().find(".rt-checkboxWrapper"),
                event = null;
            if (!$(this).hasClass("checked")) {
                $(this).addClass("checked");
                $(this).find(".rt-checkbox")[0].checked = true;
                childCheckbox.each(function() {
                    $(this).addClass("checked");
                    $(this).find(".rt-checkbox")[0].checked = true;
                });
                checkParent($(this));
                event = $.Event("nodeOnCheck");
            } else {
                $(this).removeClass("checked");
                $(this).find(".rt-checkbox")[0].checked = false;
                childCheckbox.each(function() {
                    $(this).removeClass("checked");
                    $(this).find(".rt-checkbox")[0].checked = false;
                });
                var parentNodes = $(this).parents(".rt-node").each(function() {
                    var checkbox = $(this).children(".rt-checkboxWrapper");
                    checkbox.removeClass("checked");
                    checkbox.find(".rt-checkbox")[0].checked = false;
                });
                event = $.Event("nodeOnUncheck");
            }
            $(this).trigger(event);
        };
        //导出为Excel表格函数
        var exportExcel = function() {
            window.open(serverURL + "/GenerateExcel.aspx" + getQuery() + "&ConfigFile=" + settings.configFile);
        };
        var nodeOnClick = function() {
            var elem = $(".rt-body").children(".rt-search-result");
            if (!elem.length) {
                var event = $.Event("nodeOnClick");
                $(this).trigger(event);
            }
        };

        load();
    };

    $.fn.rtGetCheckedRows = function() {
        var elemID = $(this).attr("id");
        return cachedRows[elemID];
    };

    $.fn.rtGetAllRowStr = function() {
        var allRows = [];
        $(this).find("tbody").children("tr").each(function() {
            var row = {};
            $(this).children("td").each(function() {
                var col = $(this).attr("name");
                if (col !== "rt-td-checkbox" && col !== "操作") {
                    row[col] = $(this).attr("data-value");
                }
            });
            allRows.push(JSON.stringify(row));
        });
        return allRows;
    };

    $.fn.rtAppendNewRow = function() {
        var elemID = $(this).attr("id");
        $(this).find("tbody").append(nullRows[elemID]);
    };

}(jQuery));