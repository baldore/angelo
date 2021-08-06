import { createSlice, Dispatch, PayloadAction } from '@reduxjs/toolkit'

type Notification = {
  message: string
}

const notificationsSlice = createSlice({
  initialState: [],
  name: 'notifications',
  reducers: {
    addNotification(
      state,
      { payload: notification }: PayloadAction<Notification>
    ) {
      console.log('adding notification', notification)
      return state
    },
  },
})

const { addNotification } = notificationsSlice.actions

export const displayNotification =
  (notification: Notification) => (dispatch: Dispatch) => {
    dispatch(addNotification(notification))
  }

export const reducer = notificationsSlice.reducer
