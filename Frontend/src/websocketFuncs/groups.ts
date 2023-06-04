import { useGroupsStore } from "../stores/groups";
import type { group, gchatmsg } from "../interfaces/interfaces";

export const allGroups = (groups: group[]) => {
  if (groups === null) {
    return;
  }
  useGroupsStore().setGroups(groups);
};

export const allUGroups = (uGroups: group[]) => {
  if (uGroups === null) {
    return;
  }
  useGroupsStore().setUGroups(uGroups);
};

export const oneGroup = (group: object) => {
  useGroupsStore().setGroup(group);
};

export const oneGPosts = (gPosts: object) => {
  useGroupsStore().setGPosts(gPosts);
};

export const oneGOutsiders = (outsiders: object) => {
  useGroupsStore().setOutsiders(outsiders);
};

export const oneGEvents = (events: object) => {
  useGroupsStore().setEvents(events);
};

export const oneGOwner = (owner: object) => {
  useGroupsStore().setOwner(owner);
};

export const oneGStatus = (uStatus: boolean) => {
  useGroupsStore().setUStatus(uStatus);
};

export const loadGChats = (chats: gchatmsg[]) => {
  if (useGroupsStore().gChatVal === "onegroup") {
    useGroupsStore().setGChat(chats);
  } else if (useGroupsStore().gChatVal === "newgroupchat" || useGroupsStore().gChatVal === ""){
    useGroupsStore().prependMessage(chats[chats.length - 1]);
  } else if (useGroupsStore().gChatVal === "scroll"){
    useGroupsStore().appendMessages(chats);
  }
};
