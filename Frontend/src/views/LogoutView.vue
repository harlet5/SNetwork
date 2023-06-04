<script setup lang="ts">
import { useAuthStore } from "../stores/auth";
import Ws from "@/Websocket"


/* const store = useAuthStore();
store.logOut() */
Ws.send({
    Page: "logout",
    Data: {
      cid: useAuthStore().userId,
    },
  }
);
useAuthStore().logOut()
</script>
