import { useChatsStore } from "../stores/chats";
import { useAuthStore } from "@/stores/auth";
import type { Message, Chat } from "@/interfaces/interfaces";
/*
username: 123, userid: 5
username: name, userid: 4
kui peale klikin kasutajale -> teine
saadan kasutajale yhe s6numi -> teine
kasutaja klikib minu nimele omas aknas -> esimene
kasutaja kirjutab kirja mulle -> teine

*/
export const loadMessages = (
  messages: Message[],
  otherid: string,
  unseen: number
) => {
  if (
    useChatsStore().activeChat.toString() === otherid &&
    unseen === 0 &&
    useAuthStore().userId.toString() !== otherid &&
    useChatsStore().count !== 0 &&
    useChatsStore().lastAct === "onechat"
  ) {
    if (messages === null) {
      return;
    }
    console.log("appending messages");
    useChatsStore().appendMessages(messages);
  } else if (
    useChatsStore().activeChat.toString() === otherid &&
    unseen === 0 &&
    useAuthStore().userId.toString() !== otherid &&
    useChatsStore().lastAct === "newchat"
  ) {
    console.log("prepending message");
    useChatsStore().prependMessage(messages[messages.length - 1]);
  } else if (
    useChatsStore().activeChat.toString() === otherid &&
    unseen !== 0 &&
    useAuthStore().userId.toString() !== otherid
  ) {
    console.log("prepending message");
    useChatsStore().prependMessage(messages[messages.length - 1]);
  } else if (
    useAuthStore().userId.toString() !== otherid &&
    useChatsStore().activeChat.toString() === otherid
  ) {
    useChatsStore().setMessages(messages);
    console.log("loading 10 new messages");
  }
};

export const loadChats = (chats: Chat[]) => {
  chats = chats.filter((chat) => chat.OId != useAuthStore().userId);
  useChatsStore().setChats(chats);
};

export const updateActiveUserInChat = (userid: number, toggle: boolean) => {
  useChatsStore().changeUserActive(userid, toggle);
};
