import { isEmpty } from 'lodash'

// Create WebSocket connection.
const socket = new WebSocket('ws://localhost:8088/ws')
const userId = Math.floor(Math.random() * 500)

const callbacks = []

const open = () => {
  // // Connection opened
  socket.addEventListener('open', function() {
    const msg = JSON.stringify({
      content: 'foo',
      receiver_type: 'Channel',
      senderId: userId
    })

    socket.send(msg)
  })

  // Keep connection alive by sending a message every second
  setInterval(() => {
    socket.send("")
  }, 1000)
}

const send = payload => {
  try {
    let encodedPayload = JSON.stringify(payload)
    console.log(encodedPayload)
    socket.send(encodedPayload);
  } catch(e) {
    console.log(e)
  }
}

const listen = (evtName, cbName, callback) => {
  callbacks.push({
    name: cbName,
    func: callback
  })

  socket.addEventListener(evtName, (evt) => {
    try {
      let payload = JSON.parse(evt.data)
      let cb = callbacks
        .filter(c => c.name === cbName)
        .pop()

      if (isEmpty(cb)) {
        return
      }

      cb.func(payload)
    } catch(e) {
      console.log(e)
      console.log('error while parsing the websocket payload')
    }
  })
}

const unlisten = id => {
  const idx = callbacks.findIndex(i => i.name === id)
  callbacks.splice(idx)
}

export default socket
export { listen, open, send, unlisten, userId }
