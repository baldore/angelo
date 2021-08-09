import React from 'react'
import {
  Alert,
  AlertDescription,
  AlertIcon,
  Box,
  CloseButton,
} from '@chakra-ui/react'
import { atom, useAtom } from 'jotai'

type Notification = {
  message: string
}

const notificationsAtom = atom<Notification[]>([])
export const addNotificationAtom = atom(
  null,
  (get, set, notification: Notification) => {
    set(notificationsAtom, [...get(notificationsAtom), notification])
  }
)

export default function Notifications() {
  const [notifications] = useAtom(notificationsAtom)

  return (
    <Box position="fixed" bottom="0" right="0" w="50%">
      {notifications.map(({ message }) => (
        <Alert key={message} status="success" variant="left-accent">
          <AlertIcon />
          <AlertDescription>{message}</AlertDescription>
          <CloseButton position="absolute" right="8px" top="8px" />
        </Alert>
      ))}
    </Box>
  )
}
