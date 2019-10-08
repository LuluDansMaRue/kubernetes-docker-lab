const utils = {}

utils.getCommandPayload = msg => {
  if (!msg.startsWith('/')) {
    return {
      name: null,
      args: null
    }
  }

  let args = msg.match(/\S+/g) || []

  return {
    name: args.shift().replace('/', ''),
    args: args.slice()
  };
}

utils.getContentFromInput = msg => {
  if (!msg.startsWith('/'))
    return msg

  return ''
}

export default utils
