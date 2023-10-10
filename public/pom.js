let timerInterval;
let minutes = 1;
let seconds = 0;
let isRunning = false;

const timerDisplay = document.getElementById("timer");
const startButton = document.getElementById("startButton");

function updateTimer() {
  timerDisplay.textContent = `${minutes.toString().padStart(2, "0")}:${seconds
    .toString()
    .padStart(2, "0")}`;
}

function startTimer() {
  if (!isRunning) {
    isRunning = true;
    startButton.textContent = "Pause";

    timerInterval = setInterval(async function () {
      if (minutes === 0 && seconds === 0) {
        clearInterval(timerInterval);
        isRunning = false;
        startButton.textContent = "Start";
        // Add code here to handle the break time
        // For now, let's reset the timer to 25 minutes
        minutes = 25;
        seconds = 0;

        await postPom();
        updateTimer();
      } else {
        if (seconds === 0) {
          minutes--;
          seconds = 59;
        } else {
          seconds--;
        }
        updateTimer();
      }
    }, 1000);
  } else {
    clearInterval(timerInterval);
    isRunning = false;
    startButton.textContent = "Resume";
  }
}

async function postPom() {
  const currentTime = new Date();
  const timeEarlier = new Date(currentTime.getTime() - 25 * 60000);
  console.log(timeEarlier);

  const res = await postData("/pom/api/addPom", {
    start: timeEarlier.toJSON(),
    end: currentTime.toJSON(),
    task: "test",
  });

  console.log(res);
}

async function postData(url = "", data = {}) {
  // Default options are marked with *
  const response = await fetch(url, {
    method: "POST", // *GET, POST, PUT, DELETE, etc.
    mode: "cors", // no-cors, *cors, same-origin
    cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
    credentials: "same-origin", // include, *same-origin, omit
    headers: {
      "Content-Type": "application/json",
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    redirect: "follow", // manual, *follow, error
    referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: JSON.stringify(data), // body data type must match "Content-Type" header
  });
  return response.json(); // parses JSON response into native JavaScript objects
}

startButton.addEventListener("click", startTimer);
updateTimer();
