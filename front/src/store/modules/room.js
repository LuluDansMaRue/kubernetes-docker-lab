import { listen, send, unlisten, userId } from '../../core/socket'
import { hasIn } from 'lodash'

const state = {
  rooms: [],
  selectedRoom: 'default',
  message: [
    {
      room: 'default',
      msg: []
    }
  ]
}

const getters = {
  messageForRoom(state) {
    return state.message.filter(m => m.room === state.selectedRoom).pop()
  }
}

const actions = {
  populateRoom({ commit }) {
    let payload = {
      id: '1',
      sender_id: userId.toString(),
      receiver_id: 'default',
      room_id: 'default',
      content: '',
      cmd: {
        name: 'room',
        args: ['list', '']
      }
    }

    send(payload)
    listen('message', 'populate', payload => {
      commit('room', payload)
    })
  },
  populateMessage({ commit }, payload) {
    commit('message', payload)
  }
}

const mutations = {
  room(state, payload) {
    if (hasIn(payload, 'content_data')) {
      state.rooms = payload.content_data
      unlisten('populate')
    }
  },
  message(state, payload) {
    let messageRoom = state.message.filter(m => m.room === state.selectedRoom).pop()
    if (parseInt(payload.sender_id, 10) === userId) {
      payload.me = true
    } else {
      payload.me = false
    }

    messageRoom.msg.push(payload)
  }
}

export default {
  state,
  actions,
  getters,
  mutations
}