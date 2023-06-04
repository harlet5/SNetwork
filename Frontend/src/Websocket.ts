import {
  userLogin,
  userSignUp,
  userLogOut,
  refreshFollowers,
} from "./websocketFuncs/auth";
import { loadPosts, loadComments, loadPost } from "./websocketFuncs/posts";
import {
  loadChats,
  loadMessages,
  updateActiveUserInChat,
} from "./websocketFuncs/chats";
import { userProfile } from "./websocketFuncs/profile";
import {
  allGroups,
  allUGroups,
  oneGroup,
  oneGPosts,
  oneGOutsiders,
  oneGEvents,
  oneGOwner,
  oneGStatus,
  loadGChats,
} from "./websocketFuncs/groups";
import { loadNotif } from "./websocketFuncs/notif";
import {
  setLoginError,
  setRegisterErrorUsername,
  setRegisterErrorEmail,
  setGroupError,
} from "./websocketFuncs/errors";
import { useAuthStore } from "./stores/auth";
import { useUnseenCountStore } from "./stores/unseencount";
import { useGroupsStore } from "./stores/groups";
import { useNotifStore } from "./stores/notif";
interface helperstruct1 {
  act: object;
  session: helperstruct2;
}
interface helperstruct2 {
  sessionid: string;
  userid: number;
}
let connection: WebSocket | undefined;
const getConnection = (): Promise<WebSocket> => {
  if (connection && connection.readyState < 2) {
    return Promise.resolve(connection);
  }
  return new Promise((resolve) => {
    if (window["WebSocket"]) {
      const conn = new WebSocket("ws://localhost:8000/ws");

      conn.onopen = function () {
        console.log("websocket connection open");
        const b: helperstruct2 = {
          sessionid: useAuthStore().sessionId,
          userid: useAuthStore().userId,
        };
        const a: helperstruct1 = {
          act: {
            Page: "",
            Data: "",
          },
          session: b,
        };
        if (useAuthStore().userId === -1) {
          return;
        }
        const c: helperstruct1 = {
          act: {
            Page: "setconnid",
            Data: {
              uid: useAuthStore().userId,
            },
          },
          session: b,
        };
        conn.send(JSON.stringify(a));
        conn.send(JSON.stringify(c));
      };

      conn.onmessage = async function (msg) {
        let jsonMsg = JSON.parse(msg.data);
        console.log("Frontend WS'sse tulev s6num: ", jsonMsg);
        if (jsonMsg.Act !== undefined) {
          switch (jsonMsg.Act) {
            case "login":
              userLogin(
                jsonMsg.Data.Udata.UName,
                jsonMsg.Data.Udata.UId,
                jsonMsg.Data.Udata.Follow.Followers,
                jsonMsg.Data.Sid
              );
              useUnseenCountStore().setUnseenChatCount(jsonMsg.Data.Unseen);
              break;
            case "signup":
              userSignUp();
              break;
            case "logout":
              userLogOut();
              break;
            case "threads":
              loadPosts(jsonMsg.Data.Threads);
              break;
            case "chatroom":
              loadChats(jsonMsg.Data);
              break;
            case "front":
              userProfile(
                jsonMsg.Data.UFirst,
                jsonMsg.Data.ULast,
                jsonMsg.Data.UAge,
                jsonMsg.Data.UGender,
                jsonMsg.Data.UEmail,
                jsonMsg.Data.UName,
                jsonMsg.Data.UTime,
                jsonMsg.Data.UPic,
                jsonMsg.Data.UText,
                jsonMsg.Data.UNick,
                jsonMsg.Data.UPriv,
                jsonMsg.Data.Fstatus,
                jsonMsg.Data.UThreads
              );
              break;
            case "groups":
              allGroups(jsonMsg.Data.AllGroups);
              allUGroups(jsonMsg.Data.InGroups);
              break;
            case "Notifications":
              loadNotif(jsonMsg.Data.Nfc);
              break;
            case jsonMsg.Act.match(/^user_/)?.input:
              loadChats(jsonMsg.Data.Others);
              loadMessages(
                jsonMsg.Data.Chats,
                jsonMsg.Act.substring(5),
                jsonMsg.Data.Unseen
              );
              useUnseenCountStore().setUnseenChatCount(jsonMsg.Data.Unseen);
              break;
            case jsonMsg.Act.match(/^thread_/)?.input:
              loadComments(jsonMsg.Data.Comms);
              loadPost(jsonMsg.Data.Body);
              break;
            case jsonMsg.Act.match(/^group_/)?.input:
              oneGroup(jsonMsg.Data.Group);
              oneGPosts(jsonMsg.Data.Threads);
              oneGOutsiders(jsonMsg.Data.Outsiders);
              oneGEvents(jsonMsg.Data.Events);
              oneGOwner(jsonMsg.Data.Creator);
              oneGStatus(jsonMsg.Data.Ustatus);
              //loadChats(jsonMsg.Data.Others)
              //loadMessages(jsonMsg.Data.Chats)
              break;
            case jsonMsg.Act.match(/^followrecaccepted_/)?.input:
              refreshFollowers(jsonMsg.Data.SenderFData);
              break;
            case jsonMsg.Act.match(/^groupchat_/)?.input:
              loadGChats(jsonMsg.Data);
              useGroupsStore().gChatVal = "";
              break;
            case "checkCookie":
              break;
            case "added_u":
              //new user logged in, i should update his status to online
              updateActiveUserInChat(jsonMsg.Data, true);
              break;
            case "loggedout":
              updateActiveUserInChat(jsonMsg.Data, false);
              break;
            case jsonMsg.Act.match(/^newchat_/)?.input:
              // user sended u a message with his id number && and you get case user_(HisId)
              break;
            case "unread":
              useUnseenCountStore().setUnseenChatCount(jsonMsg.Data);
              break;
            case "NfcNew":
              useNotifStore().addNotifCount();
              break;
            default:
              console.log("needs switch case for: ", jsonMsg.Act);
          }
        } else {
          switch (jsonMsg.Errrr) {
            case "Password is incorrect":
              setLoginError(jsonMsg.Errrr);
              break;
            case "Username used":
              setRegisterErrorUsername(jsonMsg.Errrr);
              break;
            case "Email used":
              setRegisterErrorEmail(jsonMsg.Errrr);
              break;
            case "Username doesn't exist":
              setLoginError(jsonMsg.Errrr);
              break;
            case "Group exists":
              setGroupError(jsonMsg.Errrr);
            default:
              console.log("Needs switch case for: ", jsonMsg);
          }
        }
      };
      resolve(conn);
    } else {
      alert("Your browser does not support WebSockets");
    }
  });
};

const Ws = {
  connect: async (): Promise<void> => {
    connection = await getConnection();
  },

  send: async (e: object): Promise<void> => {
    connection = await getConnection();
    const b: helperstruct2 = {
      sessionid: useAuthStore().sessionId,
      userid: useAuthStore().userId,
    };
    const a: helperstruct1 = {
      act: e,
      session: b,
    };

    connection.send(JSON.stringify(a));
    console.log("WS saadab backendi s6numi: ", JSON.stringify(a));
  },

  disconnect: (): void => {
    if (connection) {
      connection.close();
    }
  },
};

export default Ws;
