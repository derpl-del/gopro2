var $newdiv1 = "<div id='notif-login' class='alert alert-success alert-dismissible show' role='alert'>"
var $newdiv2 = "<button type='button' class='close' data-dismiss='alert' aria-label='Close'><span aria-hidden='true'>&times;</span></button></div>"
window.onload = function () {
    var ErrorCode = sessionStorage.getItem("ErrorCode");
    if (ErrorCode == "0000") {
        greeting = "Success";
        $("#notif-result").append($newdiv1 + greeting + $newdiv2)
        sessionStorage.setItem("ErrorCode", "");
    }
    else {
    }
    $(document).ready(function () {
        $("tr").click(function () {
            var input = this.id.substring(0, 3);
            if (input == "row") {
                id = this.id;
                var pid = document.getElementById(id).getAttributeNode("pid").value;
                var tittle = document.getElementById(id).getAttributeNode("title").value;
                var pname = document.getElementById(id).getAttributeNode("pname").value;
                var pamount = document.getElementById(id).getAttributeNode("pamount").value;
                var pquality = document.getElementById(id).getAttributeNode("pquality").value;
                var pprice = document.getElementById(id).getAttributeNode("pprice").value;
                var pcategory = document.getElementById(id).getAttributeNode("pcategory").value;
                document.getElementById("trxid").setAttribute("pid", pid);
                document.getElementById('tittle_name').value = tittle;
                document.getElementById('product_name').value = pname;
                document.getElementById('product_amount').value = pamount;
                document.getElementById('product_quality').value = pquality;
                document.getElementById('product_price').value = pprice;
                document.getElementById('product_category').value = pcategory;
            } else { }
        });
        $("#edit-apply").click(function () {
            var pid = document.getElementById("trxid").getAttributeNode("pid").value;
            var tittle = document.getElementById("tittle_name").value;
            var pname = document.getElementById("product_name").value;
            var pamount = document.getElementById("product_amount").value;
            var pquality = document.getElementById("product_quality").value;
            var pprice = document.getElementById("product_price").value;
            var pcategory = document.getElementById("product_category").value;
            var data = JSON.stringify({ "pid": pid, "tittle": tittle, "name": pname, "amount": pamount, "quality": pquality, "category": pcategory, "price": pprice });
            $.post("/EditHandle", data, function (data, status) {
                var response = jQuery.parseJSON(data);
                if (response.ErrorCode == "0000") {
                    sessionStorage.setItem("ErrorCode", response.ErrorCode);
                    document.location.reload();
                }
            })
        });

    })

}