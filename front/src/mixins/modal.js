// A reference to the modal library
let modal = {}

/**
 * Modal Proxy
 */
const ModalProxy = {
  methods: {
    /**
     * Set Modal
     * 
     * @param {Object} $modal 
     */
    setModal($modal) {
      modal = $modal
    },
    /**
     * Show
     *    Open the modal
     * 
     * @param {Object} props
     * @param {String} name
     */
    show(name, props) {
      modal.show(name, props)
    },
    /**
     * Hide
     *    Close the modal
     * @param {String} name
     */
    hide(name) {
      modal.hide(name)
    }
  }
}

export default ModalProxy
