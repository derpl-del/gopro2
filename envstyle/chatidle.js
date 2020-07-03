let conn;

$(document).ready(function () {
    $("tr").click(function () {
        $("#tbody1").remove();
        $("#menu").append("<tbody id='tbody1'></tbody>");
        var roomIdlat = document.getElementById("menu").getAttributeNode("room-id").value;
        if (!conn) {
        } else {
            conn.close(1000, "Deliberate disconnection");
        }
        var input = this.id.substring(0, 3);
        if (input == "row") {
            id = this.id;
            var roomId = document.getElementById(id).getAttributeNode("chatid").value;
            var $newdiv1 = "<tr><th scope='row'>";
            var $newdiv2 = "</th></tr>";
            var $div1 = "<tr class='table-secondary'><td scope='col'>";
            var $udiv1 = "<tr class='table-primary'><td scope='col'>";
            var $div2 = "</th><td class='text-right' scope='col'>";
            var $div3 = "</th></tr><tr><td colspan='2'>";
            var $div4 = "</td></tr>";
            let msg = document.getElementById("msg");
            var sender = document.getElementById("tbody").getAttributeNode("sen").value;

            var h;
            var m;
            var s;
            var objDiv = document.getElementById("chat-menu");
            objDiv.scrollTop = 100;

            $("#menu").attr("room-id", roomId);
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
                            if (result[1] == sender) {
                                $("#tbody1").append($udiv1 + result[1] + $div2 + result[3].substring(0, (result[3].length) - 1) + $div3 + result[2] + $div4);
                            } else {
                                $("#tbody1").append($div1 + result[1] + $div2 + result[3].substring(0, (result[3].length) - 1) + $div3 + result[2] + $div4);
                            }
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
                var today = new Date();
                h = (today.getHours() < 10 ? '0' : '') + today.getHours();
                m = (today.getMinutes() < 10 ? '0' : '') + today.getMinutes();
                s = (today.getSeconds() < 10 ? '0' : '') + today.getSeconds();
                var datechat = h + ":" + m + ":" + s;
                var roomChat = document.getElementById(id).getAttributeNode("chatid").value;
                if (!conn) {
                    return false;
                }
                if (!msg.value) {
                    return false;
                }
                var senditem = roomChat + "|" + sender + "|" + msg.value + "|" + datechat + "&";
                conn.send(senditem);
                msg.value = "";
                return false;
            });


            if (window["WebSocket"]) {
                conn = new WebSocket("ws://" + document.location.host + "/ws/" + roomId);
                conn.onclose = function (evt) {

                };

                conn.onmessage = function (evt) {
                    let messages = evt.data;
                    if (messages == "") { } else {
                        var result = appendLog(messages);
                        if (result[1] == sender) {
                            $("#tbody1").append($udiv1 + result[1] + $div2 + result[3].substring(0, (result[3].length) - 1) + $div3 + result[2] + $div4);
                        } else {
                            $("#tbody1").append($div1 + result[1] + $div2 + result[3].substring(0, (result[3].length) - 1) + $div3 + result[2] + $div4);
                        }
                    }
                };

            } else {
                let item = document.createElement("div");
                greeting = "Your browser does not support WebSockets.";
                $("#tbody1").append($newdiv1 + greeting + $newdiv2)
            }
        } else { }
    });
});