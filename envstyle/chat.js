window.onload = function () {
    let conn;
    var $newdiv1 = "<tr><th scope='row'>";
    var $newdiv2 = "</th></tr>";
    var $div1 = "<tr class='table-secondary'><td scope='col'>";
    var $div2 = "</th><td class='text-right' scope='col'>";
    var $div3 = "</th></tr><tr><td colspan='2'>";
    var $div4 = "</td></tr>";
    let msg = document.getElementById("msg");
    let sender = document.getElementById("form").getAttributeNode("sen").value;
    const params = window.location.href.split("/");
    const roomId = params[params.length - 1];
    var today = new Date();
    var h = today.getHours();
    var m = today.getMinutes();
    var s = today.getSeconds();
    var datechat = h + ":" + m + ":" + s;

    var objDiv = document.getElementById("chat-menu");
    objDiv.scrollTop = 100;


    var datareq = JSON.stringify({ "ChatID": roomId });
    $.post("/UserChatVal", datareq, function (data, status) {
        var response = jQuery.parseJSON(data);
        if (response.ErrorCode == "0000") {
            $.post("/UserGetChat", datareq, function (data, status) {
                var response = jQuery.parseJSON(data);
                var message = response.Message;
                let messages = message.split('&');
                for (let i = 0; i < (messages.length - 1); i++) {
                    var result = appendLog(messages[i]);
                    $("#tbody1").append($div1 + result[1] + $div2 + result[3] + $div3 + result[2] + $div4);
                }
            })
        }
    })

    function appendLog(item) {
        let messages = item.split('|');
        var IDs = new Object();
        for (let i = 0; i < messages.length; i++) {
            IDs[i] = messages[i];
        }
        return IDs;
    }

    $("form").submit(function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        var senditem = roomId + "|" + sender + "|" + msg.value + "|" + datechat + "&";
        conn.send(senditem);
        msg.value = "";
        return false;
    });

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws/" + roomId);
        conn.onclose = function (evt) {
            greeting = "Connection closed.";
            $("#tbody1").append($newdiv1 + greeting + $newdiv2)
        };
        conn.onmessage = function (evt) {
            let messages = evt.data.split('&');
            for (let i = 0; i < (messages.length - 1); i++) {
                var result = appendLog(messages[i]);
                $("#tbody1").append($div1 + result[1] + $div2 + result[3] + $div3 + result[2] + $div4);
            }
        };
    } else {
        let item = document.createElement("div");
        greeting = "Your browser does not support WebSockets.";
        $("#tbody1").append($newdiv1 + greeting + $newdiv2)
    }
};