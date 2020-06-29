var $newdiv1 = "<div id='notif-login' class='alert alert-warning alert-dismissible show' role='alert'>"
var $signdiv1 = "<div id='notif-signup' class='alert alert-warning alert-dismissible show' role='alert'>"
var $succdiv1 = "<div id='notif-login' class='alert alert-success alert-dismissible show' role='alert'>"
var $message = "<strong>Holy guacamole!</strong> You should check in on some of those fields below"
var $newdiv2 = "<button type='button' class='close' data-dismiss='alert' aria-label='Close'><span aria-hidden='true'>&times;</span></button></div>"
window.onload = function () {
    var ErrorCode = sessionStorage.getItem("ErrorCode");
    if (ErrorCode == "0000") {
        greeting = "Success";
        $("#notif-result").append($succdiv1 + greeting + $newdiv2)
        sessionStorage.setItem("ErrorCode", "");
    }
    else {
    }
    $(document).ready(function () {

        $("#login_send").click(function () {
            $("#notif-login").alert('close')
            var username = document.getElementById("username").value;
            var password = document.getElementById("password").value;
            var data = JSON.stringify({ "username": username, "password": password });
            $.post("/UserLoginVal", data, function (data, status) {
                var response = jQuery.parseJSON(data);
                if (response.Message == "0000") {
                    greeting = "Success";
                    $.post("/login_page", data, function (data) {
                        window.location.replace("/");
                    })
                } else if (response.Message == "0001") {
                    greeting = "Failed : Invalid Username";
                    $("#notif-result").append($newdiv1 + greeting + $newdiv2)
                } else {
                    greeting = "Failed : Invalid Pasword";
                    $("#notif-result").append($newdiv1 + greeting + $newdiv2)
                }
            })
        });

        $("#signup_send").click(function () {
            $("#notif-signup").alert('close')
            var username = document.getElementById("UsernameIn").value;
            var password = document.getElementById("PasswordIn").value;
            var datareq = JSON.stringify({ "username": username, "password": password });
            $.post("/SignLoginVal", datareq, function (data, status) {
                var response = jQuery.parseJSON(data);
                if (response.Message == "0000") {
                    $.post("/signup_page", datareq, function (result, status) {
                        sessionStorage.setItem("ErrorCode", "0000");
                        document.location.reload();
                    })
                } else if (response.Message == "0001") {
                    greeting = "Failed : Username Already Exists";
                    $("#notif-result-success").append($signdiv1 + greeting + $newdiv2)
                } else if (response.Message == "0002") {
                    greeting = "Failed : Invalid Username Value";
                    $("#notif-result-success").append($signdiv1 + greeting + $newdiv2)
                } else if (response.Message == "0003") {
                    greeting = "Failed : Invalid Pasword Value";
                    $("#notif-result-success").append($signdiv1 + greeting + $newdiv2)
                }
                else {
                    greeting = "Failed : Contact Admin";
                    $("#notif-result-success").append($signdiv1 + greeting + $newdiv2)
                }
            })
        });

        $("#login-form").keypress(function (e) {
            if (e.which == 13) {
                $("#notif-login").alert('close')
                var username = document.getElementById("username").value;
                var password = document.getElementById("password").value;
                var data = JSON.stringify({ "username": username, "password": password });
                $.post("/UserLoginVal", data, function (data, status) {
                    var response = jQuery.parseJSON(data);
                    if (response.Message == "0000") {
                        greeting = "Success";
                        $.post("/login_page", data, function (data) {
                            window.location.replace("/");
                        })
                    } else if (response.Message == "0001") {
                        greeting = "Failed : Invalid Username";
                        $("#notif-result").append($newdiv1 + greeting + $newdiv2)
                    } else {
                        greeting = "Failed : Invalid Pasword";
                        $("#notif-result").append($newdiv1 + greeting + $newdiv2)
                    }
                })
            }
        });
    });


}