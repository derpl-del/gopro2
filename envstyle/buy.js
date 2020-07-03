$(document).ready(function () {
    $("#buy-button").click(function () {
        document.getElementById("buy_product").getElementsByTagName("button")[0].removeAttribute("data-toggle");
        var total_amount = document.getElementById("amount_buy").value;
        var total_product = document.getElementById("amount_product").textContent;
        if (total_product - total_amount < 0) { }
        else {
            var h1 = document.getElementById("buy_product").getElementsByTagName("button")[0];
            var att = document.createAttribute("data-toggle");
            att.value = "modal";
            h1.setAttributeNode(att);
            document.getElementById('total_amount').value = total_amount;
            var tittle = document.getElementById("name_tittle").textContent;
            document.getElementById('tittle_name').value = tittle;
            var product_name = document.getElementById("name_product").textContent;
            document.getElementById('product_name').value = product_name;
            var product_category = document.getElementById("category_product").textContent;
            document.getElementById('product_category').value = product_category;
            var product_quality = document.getElementById("quality_product").textContent;
            document.getElementById('product_quality').value = product_quality;
            var product_price = document.getElementById("price_product").textContent;
            var total_pay = product_price * total_amount;
            var total_pay = "IDR " + total_pay;
            document.getElementById('total_pay').value = total_pay;
        }
    });
    $("#confirm-buy").click(function () {
        var trx_id = document.getElementById("trxid").getAttributeNode("pid").value;
        var owner = document.getElementById("trxid").getAttributeNode("owner").value;
        var tittle = document.getElementById("tittle_name").value;
        var product_name = document.getElementById("product_name").value;
        var product_category = document.getElementById("product_category").value;
        var product_quality = document.getElementById("product_quality").value;
        var total_pay = document.getElementById("total_pay").value;
        var total_amount = document.getElementById("total_amount").value;
        var datareq = JSON.stringify({ "pid": trx_id, "owner_in": owner, "tittle": tittle, "name": product_name, "amount_buy": total_amount, "total_pay": total_pay, "category_in": product_category, "quality_in": product_quality });
        $.post("/buy_someproduct", datareq, function (data, status) {
            document.location.reload();
        })
    });
    $("#chat-button").click(function () {
        var user2 = document.getElementById("name_tittle").getAttributeNode("owner").value;
        var user1 = document.getElementById("name_tittle").getAttributeNode("buyer").value;
        var datareq = JSON.stringify({ "User1": user1, "User2": user2 });
        $.post("/GetChatVal", datareq, function (data1, status) {
            var response = jQuery.parseJSON(data1);
            var message = response.Message;
            if (message == "0001") {
                $.post("/CreateChat", datareq, function (data2, status) {
                    var response = jQuery.parseJSON(data2);
                    var message = response.Message;
                    var url = "/chat/" + message;
                    window.location.replace(url);
                })
            } else {
                var url = "/chat/" + message;
                window.location.replace(url);
            }

        })
    });

});