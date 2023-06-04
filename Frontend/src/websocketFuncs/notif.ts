import { useNotifStore } from "../stores/notif";

export const loadNotif = (notifications: object) => {
  useNotifStore().setNotif(notifications)
}

