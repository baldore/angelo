import React from 'react'
import {
  Alert,
  AlertDescription,
  AlertIcon,
  Box,
  CloseButton,
  VStack,
} from '@chakra-ui/react'
import { atom, useAtom } from 'jotai'
import { atomWithImmer } from 'jotai/immer'

const DEFAULT_REMOVE_TIME = 3000

type NotificationStatus = 'info' | 'warning' | 'success' | 'error'

type Notification = {
  id: string
  message: string
  status: NotificationStatus
}

export const notificationsAtom = atomWithImmer<Notification[]>([])

type AddNotificationOptions = { message: string; status?: NotificationStatus }

/**
 * Adds a new notification with a default timeout.
 */
export const addNotificationAtom = atom(
  null,
  (get, set, options: AddNotificationOptions) => {
    const { message, status = 'success' } = options
    const newId = Date.now().toString()

    set(notificationsAtom, (draft) => {
      draft.push({ id: newId, message, status })
    })

    setTimeout(() => {
      const index = get(notificationsAtom).findIndex(({ id }) => id === newId)
      set(notificationsAtom, (draft) => {
        draft.splice(index, 1)
      })
    }, DEFAULT_REMOVE_TIME)
  }
)

export default function Notifications() {
  const [notifications] = useAtom(notificationsAtom)

  return (
    <VStack position="fixed" bottom="3" right="3" w="50%" spacing={2}>
      {notifications.map(({ id, message, status }) => (
        <Alert key={id} status={status} variant="left-accent">
          <AlertIcon />
          <AlertDescription>{message}</AlertDescription>
          <CloseButton position="absolute" right="8px" top="8px" />
        </Alert>
      ))}
    </VStack>
  )
}
