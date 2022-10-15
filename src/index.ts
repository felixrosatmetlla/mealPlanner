#!/usr/bin/env node

import { Client, APIErrorCode } from "@notionhq/client"

// Initializing a client
const notion = new Client({
  auth: process.env.NOTION_TOKEN,
})

const recipes = async () => await notion.databases.query({
    database_id: process.env.NOTION_REC,
    // filter: {
    //   property: "Landmark",
    //   rich_text: {
    //     contains: "Bridge",
    //   },
    // },
  })
console.log(recipes)