/**
 * Created by lixintang on 2017/5/23.
 */
var hostsId ="";//当前主机id
var nodeId = "";//当前主机nodeId
var supTable;

function drawTable() {
    if(supTable){
        supTable.destroy();
    }
    supTable = $('#advancedDataTable').DataTable({
        "processing": true,
        "serverSide": true,
        "ajax": {
            "url": "/api/v1/sysDatas/page",
            "data" : {"hostsId" : hostsId },
            "type":"POST"
        },
        "filter" : false,//检索部分是否显示
        "columns": [
            { "data": "create_time" },
            { "data": "cpu_rate" },
            { "data": "disk_rate" },
            { "data": "disk_used" },
            { "data": "disk_total" },
            { "data": "load_avg_15" },
            { "data": "memory_rate" }
        ],
        "language": {
            "processing": "处理中...",
            "lengthMenu": "显示 _MENU_ 项结果",
            "zeroRecords": "没有匹配结果",
            "info": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
            "infoEmpty": "显示第 0 至 0 项结果，共 0 项",
            "infoFiltered": "(由 _MAX_ 项结果过滤)",
            "infoPostFix": "",
            "search": "搜索:",
            "url": "",
            "emptyTable": "未搜索到数据",
            "loadingRecords": "载入中...",
            "infoThousands": ",",
            "paginate": {
                "first": "首页",
                "previous": "上页",
                "next": "下页",
                "last": "末页"
            },
        },
    });
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
            drawTable()
        });
        $("#hostsSelect").trigger("liszt:updated");
        drawTable()
    })
}

