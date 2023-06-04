import { usePostsStore } from "../stores/posts";

export const loadPosts = (posts: object) => {
  usePostsStore().setPosts(posts)
}

export const loadComments = (comments: object) => {
  usePostsStore().setComments(comments)
}
export const loadPost = (post: object) => {
  usePostsStore().setPost(post)
}