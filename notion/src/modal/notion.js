import { Client } from "@notionhq/client";
import { DEFAULT_PER_PAGE } from '../constants/constants.js';
import { getHasMoreParams } from '../utils/utils.js';

const { NOTION_API_KEY, NOTION_DATABASE_ID } = process.env;
const notion = new Client({ auth: NOTION_API_KEY });

const errorWrapper = (fn) => {
  return async (params) => {
    try {
      return await fn(params);
    } catch (error) {
      console.error("🚀 ~ file: notion.js ~ error", error)
    }
  }
}

const getDataBaseContents = async () => {
  const response = await notion.databases.query({
    database_id: NOTION_DATABASE_ID,
    filter: { property: 'status', number: { equals: 1 } },
    sorts: [{ property: 'priority', direction: 'ascending' }]
  });

  const noteCategory = response.results.map(({ properties }) => ({
    title: properties.key.title[0].plain_text,
    pageId: properties.key.title[0].mention.page.id,
    pageUrl: properties.key.title[0].href,
    priority: properties.priority.number,
    category: properties.category.rich_text[0].plain_text,
  }));

  return noteCategory;
};

/**
 * @param {Object}
 ** {block_id:string, page_size:number} --> get blocks data;
 ** {block_id: string, start_cursor: string, page_size:number} --> get more blocks data;
  @description
  ** Get blocks data,
  ** If has_children is true, that means has children of blocks;
  ** If has_more is true, that means has more blocks;
*/
const getBlocksData = async (params) => {
  if (!params.block_id) return {};

  const page_size = params?.page_size || DEFAULT_PER_PAGE;
  const response = await notion.blocks.children.list({
    page_size, ...params
  });

  return {
    hasMoreParams: getHasMoreParams(params.block_id, response, page_size),
    data: response.results
  };
};

export default [
  getDataBaseContents,
  getBlocksData
].reduce((acc, fn) => {
  return { ...acc, [fn.name]: errorWrapper(fn) }
}, {});