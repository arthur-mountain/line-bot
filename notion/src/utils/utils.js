export const getHasMoreParams = (id, response, page_size) => {
  if (!response.has_more) return null;

  return {
    block_id: id,
    start_cursor: response.next_cursor,
    page_size
  }
};
