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
    fetch(url, {
        method:"POST",
        headers: {
            "Content-Type":"application/json"
        },
        body: jsonData,
    }).then(response => {
        if(response.ok){
            return response.text();
        } else {
            alert("Username or Email already used");
            return null;
        }
    })
    .then(data => {
        if(data === null) {
            return
        }
        var cleanData = JSON.parse(data);
        redirect(cleanData.token, username, email);
    })
}

function redirect(token, name, email) {

    const url = "/dashboard";
    const url2 = "/api/auth"
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

