import{a as h}from"./instance-_3v8mRCE.js";import{d as v,n as c,O as b,f as i,H as _,c as n,t,h as o,e as l,o as f,y as d,J as g}from"./index-ydEeV_dL.js";import"./request-TC7KOJ9l.js";const k={class:"__container_tabDemo3"},w={class:"option"},C={class:"__container_iframe_container"},x=["src"],M=v({__name:"monitor",setup(I){let a=c(""),e=c(!0);b(async()=>{let s=await h({});a.value=s.data});function m(){e.value=!1,setTimeout(()=>{e.value=!0},200)}function u(){window.open(a.value,"_blank")}return(s,N)=>{const r=l("a-button"),p=l("a-spin");return f(),i("div",k,[_("div",w,[n(r,{class:"btn",onClick:m},{default:t(()=>[d(" refresh ")]),_:1}),n(r,{class:"btn",onClick:u},{default:t(()=>[d(" grafana ")]),_:1})]),n(p,{class:"spin",spinning:!o(e)},{default:t(()=>[_("div",C,[o(e)?(f(),i("iframe",{key:0,id:"grafanaIframe",src:o(a),frameborder:"0"},null,8,x)):g("",!0)])]),_:1},8,["spinning"])])}}});export{M as default};
