


function headerTest() {

    $.ajax({
        url: "/header",
        type: 'get',
        dataType: 'text',
        success: function (data) {
            GenerateHTMLCallback("SUCCESS", $('#headerResponse'));
        },
        error: function(data) {
            GenerateHTMLCallback("FAILD",$('#headerResponse'));
        }
    });
};

function dataTest() {

    var req = {
        "dataSize": parseInt($('#dataReqSize').val())
    }

    $.ajax({
        url: "/data",
        type: 'post',
        dataType: 'json',
        success: function (data) {

            var s = "Received " + data.dataString.length + " characters";

            GenerateHTMLCallback(s, $('#dataResponse'));
        },
        error: function(data) {
            GenerateHTMLCallback(JSON.stringify(data),$('#dataResponse'));
        },
        data: JSON.stringify(req)
    });
};

function GenerateHTMLCallback(data, divTag) {
    divTag.html('<p></p><textarea rows=2 col=30 readonly="true">' + data + '</textarea>')
};