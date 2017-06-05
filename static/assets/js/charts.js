/**
 * Created by lixintang on 2017/5/23.
 */
var hostsId =""

var g01 = new JustGage({
    id: "gauge01",
    value: 0,
    min: 0,
    max: 100,
    title: "CPU",
    titleFontColor : "rgba(255,255,255,.6)",
    valueFontColor:  "rgba(255,255,255,.8)"
});
var g02 = new JustGage({
    id: "gauge02",
    value : 0,
    title : "内存",
    min: 0,
    max: 100,
    gaugeWidthScale: .3,
    titleFontColor : "rgba(255,255,255,.6)",
    valueFontColor:  "rgba(255,255,255,.8)"
});

var g03 = new JustGage({
    id: "gauge03",
    value : 0,
    title : "Load",
    min: 0,
    max: 100,
    startAnimationType: 'bounce',
    refreshAnimationType: 'bounce',
    refreshAnimationTime: 2500,
    titleFontColor : "rgba(255,255,255,.6)",
    valueFontColor:  "rgba(255,255,255,.8)"
});

var g04 = new JustGage({
    id: "gauge04",
    value : 0,
    title : "硬盘",
    min: 0,
    max: 100,
    donut: true,
    titleFontColor : "rgba(255,255,255,.6)",
    valueFontColor:  "rgba(255,255,255,.8)"
});

// Generate easy-pie charts
var charts = $('.easypiechart .percentage');
charts.easyPieChart({
    animate: 2000,
    onStart: function(value) {
        $(this.el).find('span').animateNumbers(Math.floor(100*Math.random()));
    }
});

//update instance every 1 sec
window.setInterval(function() {
    refreshAllCharts(false)
}, 30000);



function refreshAllCharts(initFlag) {
    refreshMintorData();
    reDrawQPSTable(initFlag);
    reDrawBusiTable(initFlag);
}

//刷新hostsId机器的系统信息图表
function refreshMintorData() {
    $.post("/api/v1/sysDatas/newest", {"hostsId": hostsId}, function (result) {
        data = JSON.parse(result.data)
        g01.refresh(parseFloat(data.cpu_rate) * 100);
        g02.refresh(parseFloat(data.memory_rate) * 100);
        g03.refresh(parseFloat(data.disk_rate) * 100);
        g04.refresh(parseFloat(data.disk_rate) * 100);
    });
}

//更新QPS表格数据
function reDrawQPSTable(flag) {
    $.post("/api/v1/servicesDatas/list", {"offset": "0", "limit": "10","hostsId": hostsId}, function (result) {
        data = JSON.parse(result.data);
        data.reverse();
        charDatas = [];
        $.each(data,function (n, value) {
            createTime = value.CreateTime.substr(11, 5);
            charData = {};
            charData['time'] = value.CreateTime;
            charData['s'] = value.QpsSucc;
            charData['f'] = value.QpsFail;
            charDatas.push(charData)
        });
        if(flag){
            qpsChart = Morris.Line({
                element: 'line-qps',
                data: charDatas,
                xkey: 'time',
                ykeys: ['s', 'f'],
                labels: ['QPSSucc', 'QPSFail'],
                lineColors:['#16a085','#FF0066'],
                gridTextColor: "#ffffff"
            });
        }else {
            qpsChart.setData(charDatas);
        }

    })
}
//更新业务表格数据
function reDrawBusiTable(flag, svcType) {
    $.post("/api/v1/busiDatas/charList", {"offset": "0", "limit": "10", "hostsId": hostsId}, function (result) {
        datas = JSON.parse(result.data);
        ykeys = []
        labels = []
        charDatas = [];
        $.each(datas,function (key, val){
            if(val.length > 0){
                ykeys.push(key)
                labels.push(key)
                $.each(val,function (n, value) {
                    createTime = value.create_time.substr(11, 5);
                    charData = {};
                    charData['time'] = value.create_time;
                    charData[key] = value.sup_query_success;
                    charDatas.push(charData)
                });
            }
        });

        if(flag){
            busiChart = Morris.Line({
                element: 'line-busi',
                data: charDatas,
                xkey: 'time',
                ykeys: ykeys,
                labels: labels,
                lineColors:['#16a085','#FF0066'],
                gridTextColor: "#ffffff"
            });
        }else {
            busiChart.setData(charDatas);
        }

    })
}
//初始化主机下拉列表
initHostsSelect();
//初始化主机列表 select
function initHostsSelect() {
    $.post("/api/v1/hosts/p", {"offset": "0", "limit": "1000"}, function (result) {
        var options = "";
        $.each(result.data, function (n, item) {
            if (n == 0){
                hostsId = item.id
            }
            options += "<option value='" + item.id + "'>" + item.name + "</option>";
        });
        $("#hostsSelect").html(options);
        $("#hostsSelect").chosen({disable_search_threshold: 10}).change(function(){
            hostsId = $(this).val();
            refreshAllCharts(false);
        });
        $("#hostsSelect").trigger("liszt:updated");
        refreshAllCharts(true);
    })
}

