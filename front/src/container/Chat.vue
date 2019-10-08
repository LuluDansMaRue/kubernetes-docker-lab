<template>
  <div class="chat">
    <div class="chat__sidebar">
      <p>User id: {{ id }}</p>
      <br />
      <p><b>Room</b></p>
      <div class="chat__sidebar_room">
        <div
          class="room"
          v-for="room in rooms"
          v-bind:key="room"
          @click="changeRoom(room)"
        >
          <p>{{ room }}</p>
        </div>
      </div>
    </div>
    <div class="chat__content">
      <div class="chat__content__view">
        <div class="view">
          <div
            class="message"
            v-for="(msg, idx) in messages.msg"
            v-bind:key="idx"
          >
            <div class="message__me" v-if="msg.me">
              <p><span>{{msg.sender_id}}: </span> {{msg.content}}</p>
            </div>
            <div class="message__others" v-else>
              <p><span>{{msg.sender_id}}: </span> {{msg.content}}</p>
            </div>
          </div>
        </div>
      </div>
      <input class="chat__content_message__bar" v-model="message" @keyup.enter="broadcastMessage"/>
    </div>
  </div>
</template>

<script>
// Beware your eyes...
import { mapState, mapGetters } from 'vuex'
import socket, { listen, send, userId } from '../core/socket'
import utils from '../core/utils'

export default {
  data() {
    return {
      message: '',
      ready: false,
      id: userId,
      room_id: 'default'
    }
  },
  created() {
    listen('message', 'default', this.listenMsg)
  },
  computed: {
    ...mapState({
      rooms: state => state.room.rooms,
    }),
    ...mapGetters({
      messages: 'messageForRoom'
    })
  },
  methods: {
    broadcastMessage() {
      let cmd = utils.getCommandPayload(this.message);
      let payload = {
        id: '1',
        sender_id: userId.toString(),
        receiver_id: 'default',
        room_id: this.room_id,
        content: utils.getContentFromInput(this.message),
        cmd: cmd
      }

      socket.send(JSON.stringify(payload));

      if (cmd.name === 'room') {
        this.$store.dispatch('populateRoom')
      }
    },
    listenMsg(evt) {
      if (!this.ready) {
        this.$store.dispatch('populateRoom')
        this.ready = true;
      }
      
      this.$store.dispatch('populateMessage', evt)
    },
    changeRoom(id) {
      this.room_id = id

      let payload = {
        id: '1',
        sender_id: userId.toString(),
        receiver_id: 'default',
        room_id: this.room_id,
        content: '',
        cmd: {
          name: 'room',
          args: ['join', id]
        }
      }

      send(payload)
    }
  }
} 
</script>

<style lang="scss" scoped>

.chat {
  display: flex;
  height: 91vh;

  &__sidebar {
    background-color: white;
    width: 15%;
  }

  &__content {
    background-color: rgb(22, 22, 22);
    width: 100%;
    position: relative;
    overflow: hidden;

    input {
      width: 90%;
      position: absolute;
      bottom: 0;
    }
  }

  .message {
    &__me {
      p {
        color: white;
      }
    }

    &__others {
      p {
        color: #eddfd4;
      }
    }
  }

  .view {
    height: 80vh;
    overflow-y: scroll;
    scrollbar-width: none; /* For Firefox */
    -ms-overflow-style: none; /* For Internet Explorer and Edge */

    &::-webkit-scrollbar {
      width: 0px;
    }
  }
}


</style>