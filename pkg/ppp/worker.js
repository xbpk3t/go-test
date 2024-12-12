/**
 * Welcome to Cloudflare Workers! This is your first worker.
 *
 * - Run "npm run dev" in your terminal to start a development server
 * - Open a browser tab at http://localhost:8787/ to see your worker in action
 * - Run "npm run deploy" to publish your worker
 *
 * Learn more at https://developers.cloudflare.com/workers/
 */


import ppp from "./pppdata.js";

// map 一下购买力数据的列表，方便搜索
const flatppp = ppp.flatMap(category => category.countries.map(countryInfo => {
  return {
    range: category.range,
    countryCode: countryInfo.country,
    countryName: countryInfo.countryName
  }
}))

// 在购买力水平列表中找到对应国家
function findCountry(countryCode) {
  return flatppp.find(deal => deal.countryCode === countryCode)
}

// 根据购买力水平，在环境变量里找到配置的折扣信息
function getDiscount(env, range) {
  switch (range) {
    case "0.0-0.1":
      return {
        code: env.level0_1 ?? "",
        discount: parseInt(env.level0_1_discount ?? "0") ?? 0
      }
    case "0.1-0.2":
      return {
        code: env.level1_2 ?? "",
        discount: parseInt(env.level1_2_discount ?? "0") ?? 0
      }
    case "0.2-0.3":
      return {
        code: env.level2_3 ?? "",
        discount: parseInt(env.level2_3_discount ?? "0") ?? 0
      }
    case "0.3-0.4":
      return {
        code: env.level3_4 ?? "",
        discount: parseInt(env.level3_4_discount ?? "0") ?? 0
      }
    case "0.4-0.5":
      return {
        code: env.level4_5 ?? "",
        discount: parseInt(env.level4_5_discount ?? "0") ?? 0
      }
    case "0.5-0.6":
      return {
        code: env.level5_6 ?? "",
        discount: parseInt(env.level5_6_discount ?? "0") ?? 0
      }
    case "0.6-0.7":
      return {
        code: env.level6_7 ?? "",
        discount: parseInt(env.level6_7_discount ?? "0") ?? 0
      }
    case "0.7-0.8":
      return {
        code: env.level7_8 ?? "",
        discount: parseInt(env.level7_8_discount ?? "0") ?? 0
      }
    case "0.8-0.9":
      return {
        code: env.level8_9 ?? "",
        discount: parseInt(env.level8_9_discount ?? "0") ?? 0
      }
    case "0.9-1.0":
      return {
        code: env.level9_10 ?? "",
        discount: parseInt(env.level9_10_discount ?? "0") ?? 0
      }
    case "1.0-1.1":
      return {
        code: env.level10_11 ?? "",
        discount: parseInt(env.level10_11_discount ?? "0") ?? 0
      }
    case "1.1-1.2":
      return {
        code: env.level11_12 ?? "",
        discount: parseInt(env.level11_12_discount ?? "0") ?? 0
      }
    case "1.2-1.3":
      return {
        code: env.level12_13 ?? "",
        discount: parseInt(env.level12_13_discount ?? "0") ?? 0
      }
    case "1.3-1.4":
      return {
        code: env.level13_14 ?? "",
        discount: parseInt(env.level13_14_discount ?? "0") ?? 0
      }
    default:
      return {code: "", discount: 0}
  }
}

// 合并国家购买力信息+折扣信息
function mergeDiscountResult(countryPPP, discount) {
  return JSON.stringify({
    range: countryPPP.range,
    countryCode: countryPPP.countryCode,
    countryName: countryPPP.countryName,
    discountCode: discount.code,
    discount: discount.discount
  });
}

// 构造 response
function responseFor(result, code) {
  return new Response(result, {
    status: code,
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Headers": "*",
      "Access-Control-Allow-Methods": "GET, OPTIONS, POST, PUT, DELETE",
      "Access-Control-Max-Age": "0"
    }
  });
}

// ✨ 核心代码
export default {
  async fetch(request, env, ctx) {
    // 1. 获取国家编码
    const countryCode = request.cf.country
    
    // 2. 在购买力列表中找到该国家
    let countryPPP = findCountry(countryCode)
    
    // 3. 通过该国家购买力获取对应优惠信息
    let discount = getDiscount(env, countryPPP.range)
    if (countryPPP && discount) {
      // 构造结果
      let result = mergeDiscountResult(countryPPP, discount)
      // 4. 可以直接返回结果供其它服务调用
      return responseFor(result, 200)
    } else {
      return responseFor("Error", 500)
    }
    
    // 5. 或者直接 301 重定向到指定优惠链接
    // let url = env.TARGET_DOMAIN
    // if (discountCode !== undefined && discountCode.length > 0) {
    //   url = env.TARGET_DOMAIN + "?checkout%5Bdiscount_code%5D=" + discountCode
    // }
    // var response = Response.redirect(url, 301);
  },
};

