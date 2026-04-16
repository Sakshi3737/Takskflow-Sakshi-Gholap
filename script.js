const API = "http://localhost:8080";
let token = localStorage.getItem("token");
let currentProject = null;

if (token) showApp();

async function login() {
  const res = await fetch(API + "/auth/login", {
    method: "POST",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify({
      email: email.value,
      password: password.value
    })
  });

  const data = await res.json();

  if (!data.token) return alert("Login failed");

  localStorage.setItem("token", data.token);
  token = data.token;

  showApp();
}

function showApp() {
  login.style.display = "none";
  app.style.display = "block";
  loadProjects();
}

async function loadProjects() {
  const res = await fetch(API + "/projects", {
    headers: {Authorization: "Bearer " + token}
  });

  const data = await res.json();
  projects.innerHTML = "";

  data.projects.forEach(p => {
    const li = document.createElement("li");
    li.innerText = p.name;
    li.onclick = () => {
      currentProject = p.id;
      loadTasks();
    };
    projects.appendChild(li);
  });
}

async function createProject() {
  await fetch(API + "/projects", {
    method: "POST",
    headers: {
      Authorization: "Bearer " + token,
      "Content-Type": "application/json"
    },
    body: JSON.stringify({name: pname.value})
  });

  loadProjects();
}

async function createTask() {
  await fetch(API + `/projects/${currentProject}/tasks`, {
    method: "POST",
    headers: {
      Authorization: "Bearer " + token,
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      title: taskTitle.value,
      status: status.value
    })
  });

  loadTasks();
}

async function loadTasks() {
  const res = await fetch(API + `/projects/${currentProject}/tasks`, {
    headers: {Authorization: "Bearer " + token}
  });

  const data = await res.json();
  tasks.innerHTML = "";

  data.tasks.forEach(t => {
    const li = document.createElement("li");
    li.innerText = `${t.title} - ${t.status}`;
    tasks.appendChild(li);
  });
}

function logout() {
  localStorage.removeItem("token");
  location.reload();
}



