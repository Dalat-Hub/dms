import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'
import getPageTitle from '@/utils/page'
import router from '@/router'
import { ElProgress } from 'element-plus'

let asyncRouterFlag = 0

const whiteList = ['Login', 'Init']

const getRouter = async(userStore) => {
  const routerStore = useRouterStore()
  await routerStore.SetAsyncRouter()
  await userStore.GetUserInfo()
  const asyncRouters = routerStore.asyncRouters
  asyncRouters.forEach(asyncRouter => {
    router.addRoute(asyncRouter)
  })
}

async function handleKeepAlive(to) {
  if (to.matched.some(item => item.meta.keepAlive)) {
    if (to.matched && to.matched.length > 2) {
      for (let i = 1; i < to.matched.length; i++) {
        const element = to.matched[i - 1]
        if (element.name === 'layout') {
          to.matched.splice(i, 1)
          await handleKeepAlive(to)
        }

        // Waiting for loading if no on-demand loading is complete
        if (typeof element.components.default === 'function') {
          await element.components.default()
          await handleKeepAlive(to)
        }
      }
    }
  }
}

router.beforeEach(async(to, from) => {
  const userStore = useUserStore()
  to.meta.matched = [...to.matched]
  handleKeepAlive(to)
  const token = userStore.token

  // Judgment in the whitelist
  document.title = getPageTitle(to.meta.title, to)
  if (whiteList.indexOf(to.name) > -1) {
    if (token) {
      if (!asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        asyncRouterFlag++
        await getRouter(userStore)
      }

      // The token can be resolved but the user id or role id that does not exist will cause infinite calls
      if (userStore.userInfo?.authority?.defaultRouter != null) {
        return { name: userStore.userInfo.authority.defaultRouter }
      } else {
        // Force logout
        userStore.ClearStorage()
        return {
          name: 'Login',
          query: {
            redirect: document.location.hash
          }
        }
      }
    } else {
      return true
    }
  } else {
    // When not on the whitelist and logged in
    if (token) {
      // Add flag to prevent multiple access to dynamic routing and stack overflow
      if (!asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        asyncRouterFlag++
        await getRouter(userStore)
        if (userStore.token) {
          return { ...to, replace: true }
        } else {
          return {
            name: 'Login',
            query: { redirect: to.href }
          }
        }
      } else {
        if (to.matched.length) {
          return true
        } else {
          return { path: '/layout/404' }
        }
      }
    }
    // When not in the whitelist and not logged in
    if (!token) {
      return {
        name: 'Login',
        query: {
          redirect: document.location.hash
        }
      }
    }
  }
})
