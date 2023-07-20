
function openPopup() {

    const popup = document.getElementsByClassName("popup")[0];
    popup.classList.add("open-popup");

}

function closePopup() {

    const popup = document.getElementsByClassName("popup")[0];


    popup.classList.remove("open-popup");
}

function createTask() {
    const taskName = document.getElementById("taskName");
    const taskContent = document.getElementById("taskContent");
    const name = getCookie("username");
    const email = getCookie("email");
    const token = getCookie("jwt-token");
    const data = {
        username: name,
        email: email, 
        token: token, 
        taskName: taskName.value, 
        taskContent: taskContent.value
    };
    taskName.value = "";
    taskContent.value = "";
    const jsonData = JSON.stringify(data);
    const url = "/api/task";
    fetch(url, {
        method: "POST", 
        body: jsonData,
    })
    .then(response => {
        if(response.ok) {
            location.reload();
        }
    });
}

function getCookie(name) {
  const cookieString = document.cookie;
  const cookies = cookieString.split('; ');

  for (const cookie of cookies) {
    const [cookieName, cookieValue] = cookie.split('=');
    if (cookieName === name) {
      return decodeURIComponent(cookieValue);
    }
  }
}


function removeTask(event) {

    if (event.target.classList.contains('removetask')) {
         
        const card = event.target.closest('.card');

        const nameElement = card.querySelector('h3');

        const name = nameElement.textContent;
        const token = getCookie("jwt-token")
        const username = getCookie("username");
        const email = getCookie("email");
        const url = "/api/task/remove";

        const data = {
            username: username,
            email: email,
            taskname: name,
            token: token
        }; 

        const jsonData = JSON.stringify(data);
        console.log(jsonData); 
        fetch(url, {
            method: "POST", 
            body: jsonData,
        }).then(response => {
            if(response.ok) {
                location.reload();
            }
        });
    }
}

const container = document.querySelector('.card');

container.addEventListener('click', function(event) {
  if (event.target.classList.contains('removetask')) {
    const taskname = event.target.closest('.taskname');
    const cont = taskname.textContent();
    console.log(cont);
    specificDiv.style.backgroundColor = 'yellow';
  }
});
