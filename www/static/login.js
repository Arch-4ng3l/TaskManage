function login() {
    
    var email = document.getElementsByName("email")[0].value;
    var pw = document.getElementsByName("passwd")[0].value;

    console.log(email);
    console.log(pw);
    var data = {
        email: email,
        password: pw
    };
    
    var jsonData = JSON.stringify(data)
    console.log(jsonData)

    url = "/api/login";
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: jsonData

    }).then(function(response){
        if(response.ok) {
            console.log("succ");
        } else {
            console.log(response);
        }
    })
}
