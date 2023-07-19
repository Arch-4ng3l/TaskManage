function login() {
    
    var email = document.getElementsByName("email")[0].value;
    var pw = document.getElementsByName("passwd")[0].value;

    var data = {
        email: email,
        password: pw
    };
    
    var jsonData = JSON.stringify(data)

    url = "/api/login";
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: jsonData

    })
    .then(response => {
        if(response.ok) {
            return response.text();
        } else {
            alert("Invalid Credentials");
        }
    })
    .then(data => {
        var cleanData = JSON.parse(data);
        var token = cleanData.token;
        var name = cleanData.username;
        redirect(token, name, email);
    })
}

function redirect(token, name, email) {

    const url = "/dashboard";
    const url2 = "/api/auth";
    const header = new Headers(); 
    header.append("jwt-token", token);
    header.append("username", name);
    header.append("email", email);

    fetch(url2, {
        method: "GET", 
        headers: header
    }).then(response => {
        if(response.ok) {

            setCookie('username', name);
            setCookie('email', email);
            setCookie('jwt-token', token);

            window.location.replace(url);
        } else {
            return
        }
    })
};

function setCookie(name, value) {
  document.cookie = `${name}=${encodeURIComponent(value)}; path=/`;
}

