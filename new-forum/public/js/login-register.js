document.addEventListener("DOMContentLoaded", function () {
  const loginForm = document.getElementById("login-form");
  const registerForm = document.getElementById("register-form");

  updateButtonVisibility();

  loginForm.addEventListener("submit", async function (event) {
    event.preventDefault();
    await logingUser();
  });
  registerForm.addEventListener("submit", async function (event) {
    event.preventDefault();
    await registerUser();
  });
  document
    .getElementById("logout-btn")
    .addEventListener("click", async function (event) {
      event.preventDefault();
      await logoutUser();
    });
});

async function logingUser() {
  const username = document.getElementById("login-username").value;
  const password = document.getElementById("login-password").value;

  fetch("/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  })
    .then(async (response) => {
      // Add 'async' here
      if (response.status === 200) {
        alert("Login successful.");
        window.location.href = "/";
        await updateButtonVisibility(); // 'await' is now allowed here
      } else {
        alert("Login failed. Please try again.");
      }
    })
    .catch((error) => {
      console.log("Error:", error);
    });
  console.log(username, password);
}

async function registerUser() {
  const username = document.getElementById("register-username").value;
  const email = document.getElementById("register-email").value;
  const password = document.getElementById("register-password").value;
  const confirmPassword = document.getElementById(
    "register-confirm-password"
  ).value;

  if (password !== confirmPassword) {
    alert("Passwords do not match.");
    return;
  }
  fetch("/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, email, password }),
  })
    .then(async (response) => {
      if (response.status === 201) {
        console.log(username, password, email);
        alert("Registration successful.");
        window.location.href = "/";
        await updateButtonVisibility();
      } else {
        alert("Registration failed. Please try again.");
      }
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}

function logoutUser() {
  fetch("/logout", {
    method: "POST",
  })
    .then((response) => {
      if (response.status === 200) {
        alert("Logout successful.");
        window.location.href = "/";
        getCurrentUser();
        updateButtonVisibility();
      } else {
        alert("Logout failed. Please try again.");
      }
    })
    .catch((error) => {
      console.log("Error:", error);
    });
}

function isLoggedIn() {
  return fetch("/api/verify-session", {
    method: "GET",
    credentials: "same-origin",
  })
    .then((response) => {
      if (response.ok) {
        return response.status === 200;
      } else if (response.status === 401) {
        return false;
      } else {
        throw new Error("Unexpected response status");
      }
    })
    .catch((error) => {
      console.error("Error:", error);
      return false;
    });
}

function getCurrentUser() {
  fetch("/api/currentuser", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      if (response.status === 200) {
        return response.json();
      } else {
        throw new Error("Not logged in");
      }
    })
    .then((data) => {
      console.log("Current user's username:", data.username);
    })
    .catch((error) => {
      console.log("Error:", error);
    });
}

async function updateButtonVisibility() {
  const loginBtn = document.getElementById("login-btn");
  const registerBtn = document.getElementById("register-btn");
  const logoutBtn = document.getElementById("logout-btn");

  const loggedIn = await isLoggedIn();

  if (loggedIn) {
    loginBtn.style.display = "none";
    registerBtn.style.display = "none";
    logoutBtn.style.display = "inline";
  } else {
    loginBtn.style.display = "inline";
    registerBtn.style.display = "inline";
    logoutBtn.style.display = "none";
  }
}
