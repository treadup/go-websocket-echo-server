console.log("Loading script file")

const socket = new WebSocket("ws://localhost:8080/echo");

socket.onopen = function(event) {
  console.log("Opened websocket connection to server")
}

socket.onmessage = function(event) {
  console.log("Recieved message: " + event.data)
  const listNode = document.createElement("li");
  const textnode = document.createTextNode(event.data);
  listNode.appendChild(textnode);
  document.getElementById("list").appendChild(listNode);
}

socket.onerror = function(error) {
  console.log("Error: " + error.message)
}

socket.onclose = function(event) {
  console.log("Websocket connection was closed")
}

function handleClick() {
  const textbox = document.getElementById('textbox')
  const message = textbox.value
  if(message) {
    console.log("Sending message " + message)
    socket.send(message)
  } else {
    console.log("The textbox does not contain a message")
  }
}

function init() {
  console.log("Initalizing application")
  const button = document.getElementById('button')
  button.addEventListener('click', handleClick)
}

window.onload = init
