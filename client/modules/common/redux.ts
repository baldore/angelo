import { configureStore } from '@reduxjs/toolkit'
import { reducer as notificationsReducer } from 'modules/notifications/NotificationsReducer'

export const store = configureStore({
  reducer: {
    notifications: notificationsReducer,
  },
})
