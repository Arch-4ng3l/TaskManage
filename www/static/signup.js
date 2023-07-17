function send() {
    url = "/api/account";

    var username = document.getElementsByName("username")[0].value;
    var email = document.getElementsByName("email")[0].value;
    var passwd = document.getElementsByName("passwd")[0].value;
    var jsonData = JSON.stringify({
        username: username, 
        email: email, 
        password: passwd
    });
    console.log(jsonData);
    fetch(url, {
        method:"POST",
        headers: {
            "Content-Type":"application/json"
        },
        body: jsonData,
    }).then(function(response) {
        if(response.ok){
            console.log("suc");
            // Get JWT Token From Body
            GetToken(response);
        } else {
            console.log(response);
        }
    });
}

function GetToken(response) {
    return response
}
