

var HomePageChart= function() {
    // Chart.js Chart, for more examples you can check out http://www.chartjs.org/docs
    var initDashChartJS = function(){
        // Get Chart Container
        var $dashChartLinesCon  = jQuery('.js-dash-chartjs-lines')[0].getContext('2d');

        // Set Chart and Chart Data variables
        var $dashChartLines, $dashChartLinesData;
        
        // Lines Chart Data
        var $dashChartLineconfig = {
            type: 'line',
            data: {
                labels: ['9-22', '9-23', '9-24', '9-25', '9-26', '9-27', '9-28'],
                datasets: [
                {
                    label: '充值金额',
                    fillColor: 'rgba(44, 52, 63, .07)',
                    strokeColor: 'rgba(44, 52, 63, .25)',
                    pointColor: 'rgba(44, 52, 63, .25)',
                    pointStrokeColor: '#fff',
                    pointHighlightFill: '#fff',
                    pointHighlightStroke: 'rgba(44, 52, 63, 1)',
                    data: [34, 42, 40, 65, 48, 56, 80],
                },
                {
                    label: '成本',
                    fillColor: 'rgba(44, 52, 63, .1)',
                    strokeColor: 'rgba(44, 52, 63, .55)',
                    pointColor: 'rgba(44, 52, 63, .55)',
                    pointStrokeColor: '#fff',
                    pointHighlightFill: '#fff',
                    pointHighlightStroke: 'rgba(44, 52, 63, 1)',
                    display:true,
                    data: [18, 19, 20, 35, 23, 28, 50]
                },
                {
                    label: '利润',
                    fillColor: 'rgba(44, 52, 63, .1)',
                    strokeColor: 'rgba(44, 52, 63, .55)',
                    pointColor: 'rgba(44, 52, 63, .55)',
                    pointStrokeColor: '#fff',
                    pointHighlightFill: '#fff',
                    pointHighlightStroke: 'rgba(44, 52, 63, 1)',
                    data: [1, 9, 2, 5, 3, 8, 5]
                }
                ]
            },
            options: {
                responsive: true,
                title:{
                    display:true
                },
                tooltips: {
                    mode: 'index',
                    intersect: false,
                },
                hover: {
                    mode: 'nearest',
                    intersect: true
                },
                scales: {
                    xAxes: [{
                        display: true,
                        scaleLabel: {
                            display: true,
                            labelString: '日期'
                        }
                    }],
                    yAxes: [{
                        display: true,
                        scaleLabel: {
                            display: true,
                            labelString: '金额'
                        }
                    }]
                }
            }
        };
        // Init Lines Chart
        $dashChartLines = new Chart($dashChartLinesCon,$dashChartLineconfig);
    };

    return {
        init: function () {
            // Init ChartJS chart
            initDashChartJS();
        }
    };
}();

// Initialize when page loads
jQuery(function(){ HomePageChart.init(); });