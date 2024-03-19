import {addComment, getCommentList} from '~/api/comment.js';
import {ref} from 'vue';

export function useComment(videoId) {
  const Comments = ref([]);
  const page = ref(1);
  const Loading = ref(false);
  const More = ref(false);
  const createComment = async (content) => {
    await addComment(videoId, content).then((res) => {
      Comments.value.push(res.data.comment);
    });
  };

  const loadComments = () => {
    if (More.value) {
      return;
    }
    Loading.value = true;
    getCommentList(videoId, page.value, 5).then((res) => {
      if (res.data.comments.length === 0) {
        Loading.value = false;
        More.value = true;
        return;
      }
      if (page.value === 1) {
        Comments.value = res.data.comments;
      } else {
        Comments.value = Comments.value.concat(res.data.comments);
      }
      Loading.value = false;
    });
    page.value++;
  };
  return {
    Comments,
    createComment,
    loadComments,
    Loading,
    More,
  };
}
