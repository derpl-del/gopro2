function myFunction() {
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
}

function exitModel() {
    var h1 = document.getElementById("buy_product").getElementsByTagName("button")[0].removeAttribute("data-toggle");
}

function loadDoc() {
    var jml = document.getElementById("trxid").getAttributeNode("total").value;
    var total_amount = document.getElementById("total_amount").value;
    if (jml - total_amount < 0) {

    }
    else {
        let xhr = new XMLHttpRequest();
        let url = "/buy_someproduct";
        var trx_id = document.getElementById("trxid").getAttributeNode("value").value;
        var tittle = document.getElementById("tittle_name").value;
        var product_name = document.getElementById("product_name").value;
        var product_category = document.getElementById("product_category").value;
        var product_quality = document.getElementById("product_quality").value;
        var total_pay = document.getElementById("total_pay").value;
        xhr.onreadystatechange = function () {
            if (this.readyState == 4 && this.status == 200) {
                location.href = "/"
            }
        };
        xhr.open("POST", url, true);
        xhr.setRequestHeader("Content-Type", "application/json");
        var data = JSON.stringify({ "pid": trx_id, "tittle": tittle, "name": product_name, "amount_buy": total_amount, "total_pay": total_pay, "category_in": product_category, "quality_in": product_quality });
        xhr.send(data);
    }
}

function removeDoc() {
    var xhttp = new XMLHttpRequest();
    let url = "/logout";
    xhttp.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
            location.href = "/"
        }
    };
    xhttp.open("POST", url, true);
    xhttp.send();
}