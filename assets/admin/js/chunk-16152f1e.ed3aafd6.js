(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-16152f1e"],{"159b":function(t,e,n){var i=n("da84"),a=n("fdbc"),s=n("785a"),c=n("17c2"),o=n("9112"),r=function(t){if(t&&t.forEach!==c)try{o(t,"forEach",c)}catch(e){t.forEach=c}};for(var l in a)a[l]&&r(i[l]&&i[l].prototype);r(s)},"17c2":function(t,e,n){"use strict";var i=n("b727").forEach,a=n("a640"),s=a("forEach");t.exports=s?[].forEach:function(t){return i(this,t,arguments.length>1?arguments[1]:void 0)}},2055:function(t,e,n){"use strict";n("8729")},"2f53":function(t,e,n){},3330:function(t,e,n){"use strict";n("c71e")},"60b8":function(t,e,n){"use strict";n("c134")},"73cd":function(t,e,n){"use strict";n("2f53")},8729:function(t,e,n){},a434:function(t,e,n){"use strict";var i=n("23e7"),a=n("da84"),s=n("23cb"),c=n("5926"),o=n("07fa"),r=n("7b0b"),l=n("65f0"),u=n("8418"),h=n("1dde"),d=h("splice"),f=a.TypeError,m=Math.max,p=Math.min,b=9007199254740991,v="Maximum allowed length exceeded";i({target:"Array",proto:!0,forced:!d},{splice:function(t,e){var n,i,a,h,d,g,y=r(this),w=o(y),_=s(t,w),x=arguments.length;if(0===x?n=i=0:1===x?(n=0,i=w-_):(n=x-2,i=p(m(c(e),0),w-_)),w+n-i>b)throw f(v);for(a=l(y,i),h=0;h<i;h++)d=_+h,d in y&&u(a,h,y[d]);if(a.length=i,n<i){for(h=_;h<w-i;h++)d=h+i,g=h+n,d in y?y[g]=y[d]:delete y[g];for(h=w;h>w-i+n;h--)delete y[h-1]}else if(n>i)for(h=w-i;h>_;h--)d=h+i-1,g=h+n-1,d in y?y[g]=y[d]:delete y[g];for(h=0;h<n;h++)y[h+_]=arguments[h+2];return y.length=w-i+n,a}})},a640:function(t,e,n){"use strict";var i=n("d039");t.exports=function(t,e){var n=[][t];return!!n&&i((function(){n.call(null,e||function(){return 1},1)}))}},c134:function(t,e,n){},c71e:function(t,e,n){},c8c3:function(t,e,n){"use strict";var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"page-view"},[n("div",{staticClass:"page-header"},[n("BreadCrumb"),t._t("header")],2),n("t-content",{staticClass:"content-wrapper"},[t._t("content")],2)],1)},a=[],s=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"breadcrumb-container"},[n("t-breadcrumb",{attrs:{maxItemWidth:"150"}},t._l(this.breadcrumbItems,(function(e){return n("t-breadcrumbItem",{key:e.name},["Index"===e.name?n("icon",{attrs:{name:t.icon}}):t._e(),t._v(" "+t._s(e.meta.title))],1)})),1)],1)},c=[],o=(n("d3b7"),n("159b"),n("f646")),r={name:"BreadCrumb",data:function(){return{icon:""}},computed:{breadcrumbItems:function(){var t=this,e=[];return this.$route.matched.forEach((function(n){e.push(n),n.meta.icon&&(t.icon=n.meta.icon)})),e}},components:{Icon:o["a"]}},l=r,u=(n("60b8"),n("2877")),h=Object(u["a"])(l,s,c,!1,null,"0934365c",null),d=h.exports,f={name:"PageView",components:{BreadCrumb:d}},m=f,p=(n("ebf7"),Object(u["a"])(m,i,a,!1,null,"a9fb7624",null));e["a"]=p.exports},d81d:function(t,e,n){"use strict";var i=n("23e7"),a=n("b727").map,s=n("1dde"),c=s("map");i({target:"Array",proto:!0,forced:!c},{map:function(t){return a(this,t,arguments.length>1?arguments[1]:void 0)}})},e475:function(t,e,n){},ebf7:function(t,e,n){"use strict";n("e475")},f06c:function(t,e,n){"use strict";var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("t-drawer",{attrs:{header:"附件列表",visible:t.visible,onClose:t.close,closeBtn:!0,size:"400px",zIndex:2600,onOverlayClick:t.close,placement:"right"},scopedSlots:t._u([{key:"footer",fn:function(){return[n("t-button",{on:{click:t.onClickUpload}},[n("UploadIcon",{attrs:{slot:"icon"},slot:"icon"}),t._v(" 上传附件 ")],1),"multiple"===t.mode?n("span",{staticClass:"complete-button"},[t.selectedAttaches.length>0?n("t-button",{attrs:{theme:"danger"},on:{click:t.onClickComplete}},[t._v(" 完成 ")]):n("t-button",{attrs:{theme:"danger",disabled:""}},[t._v("完成")])],1):t._e()]},proxy:!0}])},[n("div",{staticClass:"images-wrapper"},t._l(t.attachments,(function(e,i){return n("div",{key:e.id,staticClass:"image-item",style:t.selectedIndex[i]?{border:"5px solid #1890ff"}:{},on:{click:function(n){return t.handleSelect(e,i)},mouseenter:function(e){return t.mouseEnter(i)},mouseleave:function(e){return t.mouseLeave(i)}}},[n("img",{attrs:{src:e.url}}),n("div",{directives:[{name:"show",rawName:"v-show",value:t.maskVisible[i],expression:"maskVisible[index]"}],staticClass:"image-item-mask"},[n("div",{on:{click:function(n){return n.stopPropagation(),t.handlePreview(e)}}},[n("BrowseIcon",{staticStyle:{cursor:"pointer"}})],1),n("t-divider",{attrs:{layout:"vertical"}}),n("div",{on:{click:function(n){return n.stopPropagation(),t.handleDetail(e)}}},[n("InfoCircleIcon",{staticStyle:{cursor:"pointer"}})],1)],1)])})),0)]),n("AttachDetailDrawer",{ref:"attachDetailDrawer"}),n("AttachUploadDialog",{ref:"attachUploadDialog",on:{uploadSuccess:t.syncData}})],1)},a=[],s=(n("d81d"),n("a434"),function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("t-dialog",{attrs:{visible:t.visible,header:null,confirmBtn:null,cancelBtn:null,onClose:t.close,zIndex:2601}},[n("div",{staticClass:"upload-container"},[n("t-upload",{attrs:{theme:"image",onSuccess:t.handleUploadSuccess,onFail:t.handleUploadFail,requestMethod:t.uploadMethod,draggable:"",accept:"image/*"},model:{value:t.files,callback:function(e){t.files=e},expression:"files"}})],1)])}),c=[],o=(n("d3b7"),{name:"AttachUploadDialog",data:function(){return{visible:!1,files:[]}},computed:{},methods:{open:function(){this.visible=!0},close:function(){this.visible=!1},handleUploadFail:function(){this.$message.error("上传失败")},handleUploadSuccess:function(){this.$message.success("上传成功"),this.$emit("uploadSuccess"),this.files=[],this.close()},uploadMethod:function(t){var e=this;return new Promise((function(n){t.percent=0;var i=new FormData;i.append("file",t.raw),e.$request.post("upload",i,{headers:{"Content-Type":"multipart/form-data"}}).then((function(t){0===t.status&&n({status:"success",response:{url:t.data.url}})})).catch((function(){n({status:"fail",error:"上传失败，请检查文件是否符合规范"})}))}))}}}),r=o,l=(n("2055"),n("2877")),u=Object(l["a"])(r,s,c,!1,null,"65d582b0",null),h=u.exports,d=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("t-drawer",{attrs:{header:"附件详情",visible:t.visible,onClose:t.close,closeBtn:!0,size:"480px",zIndex:2600,onOverlayClick:t.close,placement:"right"},scopedSlots:t._u([{key:"footer",fn:function(){return[n("t-button",{on:{click:t.deleteAttach}},[t._v("删除")])]},proxy:!0}])},[t.attachment?n("t-list",{attrs:{split:!0}},[n("t-list-item",[n("t-list-item-meta",{attrs:{title:"附件名"}}),t._v(" "+t._s(t.attachment.fileName)+" ")],1),n("t-list-item",[n("t-list-item-meta",{attrs:{title:"存储位置"}}),t._v(" "+t._s(t.attachmentStorageType)+" ")],1),n("t-list-item",[n("t-list-item-meta",{attrs:{title:"图片大小"}}),t._v(" "+t._s(t.attachment.size)+" ")],1),n("t-list-item",[n("t-list-item-meta",{attrs:{title:"图片尺寸"}}),t._v(" "+t._s(t.attachment.height)+" x "+t._s(t.attachment.width)+" ")],1),n("t-list-item",[n("t-list-item-meta",{attrs:{title:"上传日期"}}),t._v(" "+t._s(t.attachment.uploadTime)+" ")],1),n("t-list-item",{scopedSlots:t._u([{key:"action",fn:function(){return[n("t-button",{attrs:{variant:"text",shape:"square"},on:{click:t.copyURL}},[n("file-copy-icon")],1)]},proxy:!0}],null,!1,320530126)},[n("t-list-item-meta",{attrs:{title:"图片链接"}}),t._v(" "+t._s(t.attachment.url)+" ")],1),n("t-list-item",{scopedSlots:t._u([{key:"action",fn:function(){return[n("t-button",{attrs:{variant:"text",shape:"square"},on:{click:t.copyMarkdown}},[n("file-copy-icon")],1)]},proxy:!0}],null,!1,765174242)},[n("t-list-item-meta",{attrs:{title:"Markdown 链接"}}),t._v(" !["+t._s(t.attachment.fileName)+"]("+t._s(t.attachment.url)+") ")],1)],1):t._e()],1)],1)},f=[],m=(n("99af"),n("4630")),p={name:"AttachDetailDrawer",data:function(){return{attachment:{},visible:!1}},computed:{attachmentStorageType:function(){switch(this.attachment.type){case 0:return"本地";default:return"未知"}}},methods:{open:function(t){this.visible=!0,this.attachment=t},close:function(){this.visible=!1},deleteAttach:function(){this.$message.info("未实现")},copyURL:function(){this.$copyText(this.attachment.url),this.$message.success("已复制")},copyMarkdown:function(){this.$copyText("![".concat(this.attachment.fileName,"](").concat(this.attachment.url,")")),this.$message.success("已复制")}},components:{FileCopyIcon:m["a"]}},b=p,v=(n("3330"),Object(l["a"])(b,d,f,!1,null,"6fcb340a",null)),g=v.exports,y=n("87eb"),w=n("c09c"),_=n("2b0e"),x=n("b25f"),k=n("6be3"),C=(n("4d26"),n("febf"),["size"]);function O(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(t);e&&(i=i.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,i)}return n}function I(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?O(Object(n),!0).forEach((function(e){Object(y["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):O(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}var D={tag:"svg",attrs:{fill:"none",viewBox:"0 0 16 16",width:"1em",height:"1em"},children:[{tag:"path",attrs:{fill:"currentColor",d:"M7.5 12V6.5h1V12h-1zM8.6 4H7.4v1.2h1.2V4z",fillOpacity:.9}},{tag:"path",attrs:{fill:"currentColor",d:"M1 8a7 7 0 1014 0A7 7 0 001 8zm1 0a6 6 0 1112 0A6 6 0 012 8z",fillOpacity:.9}}]},$=_["default"].extend({name:"InfoCircleIcon",functional:!0,props:{size:{type:String},onClick:{type:Function}},render:function(t,e){var n=e.props,i=e.data,a=n.size,s=Object(w["a"])(n,C),c=Object(k["a"])(a),o=c.className,r=c.style,l=I(I({},s||{}),{},{id:"info-circle",icon:D,staticClass:o,style:r});return i.props=l,t(x["a"],i)}}),A=n("4abd"),S=n("abb0"),j={name:"AttachSelectDrawer",data:function(){return{num:0,visible:!1,attachments:[],maskVisible:[],selectedIndex:[],selectedAttaches:[]}},props:{mode:{type:String,required:!1,default:"single"}},methods:{open:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:0;this.num=t,this.visible=!0},close:function(){this.visible=!1,this.selectedIndex=this.selectedIndex.map((function(){return!1})),this.selectedAttaches=[]},onClickUpload:function(){this.$refs.attachUploadDialog.open()},syncData:function(){var t=this;this.$request.get("attachments").then((function(e){0===e.status&&(t.attachments=e.data.list?e.data.list:[]),t.maskVisible=t.attachments.map((function(){return!1})),t.selectedIndex=t.attachments.map((function(){return!1}))})).catch((function(){t.$message.error("获取附件列表失败")}))},mouseEnter:function(t){this.$set(this.maskVisible,t,!0)},mouseLeave:function(t){this.$set(this.maskVisible,t,!1)},handleDetail:function(t){this.$refs.attachDetailDrawer.open(t)},handlePreview:function(t){this.$viewerApi({images:[t.url]})},handleSelect:function(t,e){"single"===this.mode&&(this.$emit("select",t),this.close()),0!=this.num&&this.num<=this.selectedAttaches.length?this.$message.warning("最多只能选择"+this.num+"个附件"):(this.$set(this.selectedIndex,e,!this.selectedIndex[e]),this.selectedIndex[e]?this.selectedAttaches.push(t):this.selectedAttaches.splice(this.selectedAttaches.indexOf(t),1))},onClickComplete:function(){this.$emit("select",this.selectedAttaches),this.close()}},mounted:function(){this.syncData()},components:{AttachUploadDialog:h,AttachDetailDrawer:g,InfoCircleIcon:$,BrowseIcon:A["a"],UploadIcon:S["a"]}},E=j,U=(n("73cd"),Object(l["a"])(E,i,a,!1,null,"12cc894a",null));e["a"]=U.exports}}]);
//# sourceMappingURL=chunk-16152f1e.ed3aafd6.js.map