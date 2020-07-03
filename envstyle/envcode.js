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

function SendDoc() {
    var username = document.getElementById("UsernameIn").value;
    var password = document.getElementById("PasswordIn").value;
    var xhttp = new XMLHttpRequest();
    let url = "/signup_page";
    xhttp.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
            location.href = "/login"
        }
    };
    xhttp.open("POST", url, true);
    xhttp.setRequestHeader("Content-Type", "application/json");
    var data = JSON.stringify({ "username": username, "password": password });
    xhttp.send(data);
}