import { createStore, createLogger } from 'vuex'
import Persistent from './plugins/persistent'
const debug = import.meta.env.MODE !== 'production'
const files = import.meta.glob('./modules/*.js', { eager: true })

const modules = {}
Object.keys(files).forEach((c) => {
  const module = files[c].default
  const moduleName = c.replace(/^\.\/(.*)\/(.*)\.\w+$/, '$2')
  modules[moduleName] = module
})

const persistent = Persistent({
  key: 'vuex',
  modules,
  modulesKeys: {
    local: Object.keys(modules),
    session: []
  }
})

export default createStore({
  modules: {
    ...modules
  },
  strict: debug,
  plugins: debug ? [createLogger(), persistent] : [persistent]
})
