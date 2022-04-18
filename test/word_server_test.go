package test

import (
	"github.com/go-ego/gse"
	"search-trace-server/server"
	"testing"
)

const testTxt = "博客开发者文库课程问答社区插件认证开源搜索登录/注册会员中心足迹动态vartoolbarSearchExt='{\"landingWord\":[\"结巴分词性能\"],\"queryWord\":\"\",\"tag\":[\"nlp\"],\"title\":\"中文分词性能对比\"}';(function(){varbp=document.createElement('script');varcurProtocol=window.location.protocol.split(':')[0];if(curProtocol==='https'){bp.src='https://zz.bdstatic.com/linksubmit/push.js';}else{bp.src='http://push.zhanzhang.baidu.com/push.js';}vars=document.getElementsByTagName(\"script\")[0];s.parentNode.insertBefore(bp,s);})();vararticleId=85636042;varcommentscount=2;varcurentUrl=\"https://blog.csdn.net/Mr_Yener/article/details/85636042\";varmyUrl=\"https://my.csdn.net/\";varhighlight=[\"结巴分词\",\"中文分词\",\"nlp\",\"结巴\",\"分词\",\"中文\",\"性能\",\"对比\"];//高亮数组varisRecommendModule=true;varisBaiduPre=true;varbaiduCount=2;varshare_card_url=\"https://app-blog.csdn.net/share?article_id=85636042&username=Mr_Yener\"vararticleType=1;varbaiduKey=\"结巴分词性能\";varuserNewReport=true;varneedInsertBaidu=true;varrecommendRegularDomainArr=[\"blog.csdn.net/.+/article/details/\",\"download.csdn.net/download/\",\"edu.csdn.net/course/detail/\",\"ask.csdn.net/questions/\",\"bbs.csdn.net/topics/\",\"www.csdn.net/gather_.+/\"]varcodeStyle=\"\";varbaiduSearchType=\"baidulandingword\";varcanRead=true;varblogMoveHomeArticle=false;varshowPcWindowAd=false;varshowHeadWord=true;varshowSearchText=\"\";varlinkPage=true;vararticleSource=1;vararticleReport='{\"pid\":\"blog\",\"spm\":\"1001.2101\"}';varisShowToQuestion=false;varbaiduSearchChannel='pc_relevant'varbaiduSearchIdentification='.pc_relevant_paycolumn_v3'vardistRequestId='1650172540257_92599'varinitRewardObject={giver:currentUserName,anchor:username,articleId:articleId,sign:''}varisLikeStatus=false;varisUnLikeStatus=false;varstudyLearnWord=\"\";varisCurrentUserVip=false;varcommentIsNewBeat=\"control\"if(!isCookieConcision){$('.main_father').removeClass('mainfather-concision')$('.main_father.container').removeClass('container-concision')}functiongetQueryString(name){varreg=newRegExp(\"(^|&)\"+name+\"=([^&]*)(&|$)\");//构造一个含有目标参数的正则表达式对象varr=window.location.search.substr(1).match(reg);//匹配目标参数if(r!=null)returndecodeURIComponent(r[2]);return'';}functionstripscript(s){varpattern=newRegExp(\"[`~!@#$^&*()=|{}':;',\\\\[\\\\].<>/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？%]\")varrs=\"\";for(vari=0;i<s.length;i++){rs=rs+s.substr(i,1).replace(pattern,'');}returnrs;}varblogHotWords=stripscript(getQueryString('utm_term')).length>1?stripscript(getQueryString('utm_term')):''中文分词性能对比Yener丶于 2019-01-0219:58:12 发布1663收藏2分类专栏：Linux编程人工智能知识图谱文章标签：nlp版权声明：本文为博主原创文章，遵循CC4.0BY-SA版权协议，转载请附上原文出处链接和本声明。本文链接：https://blog.csdn.net/Mr_Yener/article/details/85636042版权Linux编程同时被3个专栏收录2篇文章0订阅订阅专栏人工智能5篇文章0订阅订阅专栏知识图谱4篇文章1订阅订阅专栏中文分词性能对比以下分词工具均能在Python环境中直接调用（排名不分先后）。jieba（结巴分词）免费使用HanLP（汉语言处理包）免费使用SnowNLP（中文的类库）免费使用FoolNLTK（中文处理工具包）免费使用Jiagu（甲骨NLP）免费使用pyltp（哈工大语言云）商用需要付费THULAC（清华中文词法分析工具包）商用需要付费NLPIR（汉语分词系统）付费使用分词工具使用方式见Tutorial。每个分词工具可能有多种分词模式，本文所对比的为默认分词模式。计算公式：使用方式：pipinstalljiebapipinstallpyhanlppipinstallsnownlppipinstallfoolnltkpip3installjiagupipinstallpyltp#需要下载模型，模型地址见上诉Tutorial教程，模型所放的路径与Tutorial同pipinstallthulacpipinstallpynlpir#安装好后需要使用证书，见上诉Tutorial教程gitclonehttps://github.com/ownthink/evaluation.gitcdevaluationpython3test.py#测试工具是否能够使用python3eval.py#测试（时间较长）测试集：msrpkuothermsr测试结果pku测试结果other测试结果时间测试结果$(function(){setTimeout(function(){varmathcodeList=document.querySelectorAll('.htmledit_viewsimg.mathcode');if(mathcodeList.length>0){vartestImg=newImage();testImg.onerror=function(){mathcodeList.forEach(function(item){$(item).before('<spanclass=\"img-codecogs\">\\\\('+item.alt+'\\\\)</span>');$(item).remove();})MathJax.Hub.Queue([\"Typeset\",MathJax.Hub]);}testImg.src=mathcodeList[0].src;}},1000)})Yener丶关注关注2点赞踩2评论2收藏打赏扫一扫，分享内容点击复制链接专栏目录Python-各大中文分词性能评测08-10jieba（结巴分词）免费使用HanLP（汉语言处理包）免费使用SnowNLP（中文的类库）免费使用FoolNLTK（中文处理工具包）免费使用Jiagu（甲骨NLP）免费使用pyltp（哈工大语言云）商用需要付费THULAC（清华中文词法分析工具包）商用需要付费NLPIR（汉语分词系统）付费使用中文分词常见方法m0_37710823的博客07-251万+作者：竹间智能Emotibot链接：https://www.zhihu.com/question/19578687/answer/190569700来源：知乎著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。中文分词是中文文本处理的一个基础步骤，也是中文人机自然语言交互的基础模块。不同于英文的是，中文句子中没有词的界限，因此在进行中文自然语言处理时，通常评论 2您还未登录，请先登录后发表或查看评论NLPIR分词准确率接近98.23%weixin_34010949的博客03-28148http://www.nlpir.org/几个例子：为人民办公益为/p人民/n办/v公益/n独立自主和平等互利的原则独立自主/vl和/cc平等互利/vl的/ude1原则/n结婚的和尚未结婚的结婚/vi的/ude1和/cc尚未/d结婚/vi的/ude1北京大学生前来应聘北京/ns大学生/...jieba结巴分词--关键词抽取_中文分词性能对比最新发布weixin_39709367的博客11-27138中文分词性能对比以下分词工具均能在Python环境中直接调用（排名不分先后）。jieba（结巴分词）免费使用HanLP（汉语言处理包）免费使用SnowNLP（中文的类库）免费使用FoolNLTK（中文处理工具包）免费使用Jiagu（甲骨NLP）免费使用pyltp（哈工大语言云）商用需要付费THULAC（清华中文词法分析工具包）商用需要付费NLPIR（汉语分词系统）付费使用分词工具使...四款python中文分词系统简单测试eric88的专栏12-169671四款python中文分词系统简单测试：注：中科院分词可采用调用C库的方式使用纠正下：中科院分词2012支持关键词提取准确率测试（使用对应项目提供在线测试，未添加用户自定义词典）结巴中文分词http://209.222.69.242:9000/中科院分词系统http://ictclas.org/ictclas_demo.html中文分词器性能比较06-071万+原文：http://www.cnblogs.com/wgp13x/p/3748764.html摘要：本篇是本人在Solr的基础上，配置了中文分词器，并对其进行的性能测试总结，具体包括使用mmseg4j、IKAnalyzer、Ansj，分别从创建索引效果、创建索引性能、数据搜索效率等方面进行衡量。 具体的Solr使用方法假设读者已有了基础，关于SoNLP+词法系列（一）︱中文分词技术小结、几大分词引擎的介绍与比较素质云笔记11-254万+笔者想说：觉得英文与中文分词有很大的区别，毕竟中文的表达方式跟英语有很大区别，而且语言组合形式丰富，如果把国外的内容强行搬过来用，不一样是最好的。所以这边看到有几家大牛都在中文分词以及NLP上越走越远。哈工大以及北大的张华平教授（NLPIR）的研究成果非常棒！但是商业应用的过程中存在的以下的问题：1、是否先利用开源的分词平台进行分词后，再自己写一些算法进行未登录词、歧义词的识别？2、或中文分词器比较仗贱走天涯02-139362http://blog.csdn.net/u013063153/article/details/72904322结巴分词java高性能实现，优雅易用的api设计，性能优于huabanjieba分词01-141万+SegmentSegment是基于结巴分词词库实现的更加灵活，高性能的java分词实现。变更日志创作目的分词是做NLP相关工作，非常基础的一项功能。jieba-analysis作为一款非常受欢迎的分词实现，个人实现的opencc4j之前一直使用其作为分词。但是随着对分词的了解，发现结巴分词对于一些配置上不够灵活。有很多功能无法指定关闭，比如HMM对于繁简体转换...主流分词工具性能测试结果对比木东的博客08-282275JiebaTimeexpenditure:159.01109504699707Accuracyscore:0.8949003450398567Macrof1score:0.8539702787662644Microf1score:0.8949003450398567Classificationreport:       precision  re...python︱六款中文分词模块尝试:jieba、THULAC、SnowNLP、pynlpir、CoreNLP、pyLTP热门推荐素质云笔记08-107万+公众号“素质云笔记”定期更新博客内容：THULAC四款python中中文分词的尝试。尝试的有：jieba、SnowNLP（MIT）、pynlpir（大数据搜索挖掘实验室（北京市海量语言信息处理与云计算应用工程技术研究中心））、thulac（清华大学自然语言处理与社会人文计算实验室）四款都有分词功能，本博客只介绍作者比较感兴趣、每个模块的内容。jieba在这不做介绍，可见博客：...python的中文文本挖掘库snownlp进行购物评论文本情感分析实例yyxyyx10的博客03-165万+昨晚上发现了snownlp这个库，欣喜若狂。先说说我这么开心的原因。我本科毕业设计做的是文本挖掘，用R语言做的，发现R语言对文本处理特别不友好，没有很多强大的库，特别是针对中文文本的，加上那时候还没有学机器学习算法。所以很头疼，后来不得已用了一个可视化的软件RostCM，但是一般可视化软件最大的缺点是无法调参，很死板，准确率并不高。现在研一，机器学习算法学完以后，又想起来要继续学习文本挖掘了。所以python+snownlp正负面分析大海里的金鱼的博客06-297372正负面分析背景：需要对新闻，评论做正负面分析步骤：1.安装snownlppipinstallsnownlp2.训练或导入模型训练fromsnownlpimportSnowNLP#加载情感分析模块fromsnownlpimportsentimenttext='大麦多开一个口ok？？？？正在现场俩口闲死​​​​'#文本s=SnowNLP...6种分词工具的效率、效果对比nameforcsdn的博客07-053843转自：https://www.jianshu.com/p/575fd73ce379六种分词器使用建议：对命名实体识别要求较高的可以选择HanLP，根据说明其训练的语料比较多，载入了很多实体库，通过测试在实体边界的识别上有一定的优势。中科院的分词，是学术界比较权威的，对比来看哈工大的分词器也具有比较高的优势。同时这两款分词器的安装虽然不难，但比较jieba的安装显得繁琐一点，代码迁移性会相对弱一点...中文分词效果对比weixin_33862041的博客04-291892019独角兽企业重金招聘Python工程师标准>>>...HanLPvsLTP分词功能测试weixin_33980459的博客04-19154文章摘自github,本次测试选用HanLP1.6.0,LTP3.4.0测试思路使用同一份语料训练两个分词库，同一份测试数据测试两个分词库的性能。语料库选取1998年01月的人民日报语料库。199801人民日报语料该词库带有词性标注，为了遵循LTP的训练数据集格式，需要处理掉词性标注。测试数据选择SIGHan2005提供的开放测试集。SIGHan200...中文分词模型-pkuseg-jieba-thulac对比小白水手的博客12-06913下载了4个模型，官网也有其他的模型可参考。看下模型结果对比：importpkusegs=\"小米粒儿\"seg=pkuseg.pkuseg(model_name='./web')text=seg.cut(s)print(text)'''['小米','粒儿']'''importpkusegs=\"小米粒儿\"seg=pkuseg.pkuseg(model_na...舆情分析-snownlp实战Zenbo评论分析vivian的专栏03-091157环境python2.7+pycharm,windows环境python已经抓取了评论数据情感分析思路，先将每句话分句，然后对每个短句做情感分析。因为评论有些会很长，如果整句做分析，识别率会降低。Code：#-*-coding:utf-8-*-\"\"\"SnowNLPforzenbocomments\"\"\"fromsnownlpimport...中文分词评价指标tianya111cy的专栏11-151008准确率(Precision)和召回率(Recall)Precision=正确切分出的词的数目/切分出的词的总数Recall=正确切分出的词的数目/应切分出的词的总数 综合性能指标F-measureFβ=(β2+1)*Precision*Recall/(β2*Precision+Recall)β为权重因子，如果将准确率和召回率同等看待，取β=1，就得到...“相关推荐”对你有帮助么？非常没帮助没帮助一般有帮助非常有帮助提交©️2022CSDN皮肤主题：大白设计师：CSDN官方博客返回首页关于我们招贤纳士商务合作寻求报道400-660-0108kefu@csdn.net在线客服工作时间 8:30-22:00公安备案号11010502030143京ICP备19004658号京网文〔2020〕1039-165号经营性网站备案信息北京互联网违法和不良信息举报中心家长监护网络110报警服务中国互联网举报中心Chrome商店下载©1999-2022北京创新乐知网络技术有限公司版权与免责声明版权申诉出版物许可证营业执照window.csdn.csdnFooter.options={el:'.blog-footer-bottom',type:2}Yener丶CSDN认证博客专家CSDN认证企业博客码龄7年暂无认证5原创25万+周排名179万+总排名4万+访问等级250积分20粉丝11获赞15评论73收藏私信关注热门文章知识图谱可视化展示32596知识图谱数据下载2420中文分词性能对比1663常用分词工具使用教程1410史上最大规模1.4亿中文知识图谱开源下载1409分类专栏C语言C++数据结构Python1篇Linux编程2篇Linux内核Oracle/MySQLcocos2d-xOpenCVMFCQTC#JavaAndroid人工智能5篇机器学习2篇深度学习面试题/算法服务器快速建站我的白日梦知识图谱4篇聊天机器人1篇对话机器人最新评论知识图谱数据下载weixin_40010273:下载下来怎么是csv文件知识图谱数据下载fkyyly:数据从哪里可以下载啊知识图谱数据下载造车的小同学:打不开网址啊知识图谱可视化展示长着翅膀的乌龟:谢谢大佬效果很好不过我想改改这里面的函数都没见过，能不能推荐一些教程具体讲这些函数的作用，或者你这代码有没有注释的。史上最大规模1.4亿中文知识图谱开源下载妹特斯邦威回复Yener丶:不好意思，是我弄错了，昨天没太注意看，你的发表时间更早，向你致歉。建议你去这个网址看看，有人搬运你的文章，却写的原创，和你的文章是一样的您愿意向朋友推荐“博客详情页”吗？强烈不推荐不推荐一般般推荐强烈推荐提交最新文章史上最大规模1.4亿中文知识图谱开源下载知识图谱数据下载常用分词工具使用教程2019年4篇2018年1篇(adsbygoogle=window.adsbygoogle||[]).push({});ins.adsbygoogle>.wwads-cn{display:none!important;}ins.adsbygoogle[data-ad-status='unfilled']{height:auto!important;}ins.adsbygoogle[data-ad-status='unfilled']>*{display:none!important;}ins.adsbygoogle[data-ad-status='unfilled']>.wwads-cn{display:block!important;text-align:center;/*万维广告局中*/}目录中文分词性能对比计算公式：使用方式：测试集：$(\"a.flexible-btn\").click(function(){$(this).parents('div.aside-box').removeClass('flexible-box');$(this).parents(\"p.text-center\").remove();})目录中文分词性能对比计算公式：使用方式：测试集：(adsbygoogle=window.adsbygoogle||[]).push({});ins.adsbygoogle>.wwads-cn{display:none!important;}ins.adsbygoogle[data-ad-status='unfilled']{height:auto!important;}ins.adsbygoogle[data-ad-status='unfilled']>*{display:none!important;}ins.adsbygoogle[data-ad-status='unfilled']>.wwads-cn{display:block!important;text-align:center;/*万维广告局中*/}分类专栏C语言C++数据结构Python1篇Linux编程2篇Linux内核Oracle/MySQLcocos2d-xOpenCVMFCQTC#JavaAndroid人工智能5篇机器学习2篇深度学习面试题/算法服务器快速建站我的白日梦知识图谱4篇聊天机器人1篇对话机器人vartimert=setInterval(function(){sideToolbar=$(\".csdn-side-toolbar\");if(sideToolbar.length>0){sideToolbar.css('cssText','bottom:64px!important;')clearInterval(timert);}},200);打赏作者Yener丶你的鼓励将是我创作的最大动力¥2¥4¥6¥10¥20输入1-500的整数余额支付(余额：--)扫码支付扫码支付：¥2获取中扫码支付您的余额不足，请更换扫码支付或充值打赏作者实付元使用余额支付点击重新获取扫码支付钱包余额0抵扣说明：1.余额是钱包充值的虚拟货币，按照1:1的比例进行支付金额的抵扣。2.余额无法直接购买下载，可以购买VIP、C币套餐、付费专栏及课程。余额充值//全局声明if(window.csdn===undefined){window.csdn={};}window.csdn.sideToolbar={options:{report:{isShow:true,},qr:{isShow:false,},guide:{isShow:true}}}$(function(){$(document).on('click',\"a.option-box[data-type='report']\",function(){window.csdn.userLogin.loadAjax(function(res){window.csdn.feedback({\"type\":'blog',\"rtype\":'article',\"rid\":articleId,\"reportedName\":username,\"submitOptions\":{\"title\":articleTitle,\"contentUrl\":articleDetailUrl},\"callback\":function(){showToast({text:\"感谢您的举报，我们会尽快审核！\",bottom:'10%',zindex:9000,speed:500,time:1500});}})})});})$(\".MathJax\").remove();if($('div.markdown_viewspre.prettyprintcode.hljs').length>0){$('div.markdown_views')[0].className='markdown_views';}MathJax.Hub.Config({\"HTML-CSS\":{linebreaks:{automatic:true,width:\"94%container\"},imageFont:null},tex2jax:{preview:\"none\",ignoreClass:\"title-article\"},mml2jax:{preview:'none'}});确定取消举报选择你想要举报的内容（必选）内容涉黄政治相关内容抄袭涉嫌广告内容侵权侮辱谩骂样式问题其他原文链接（必填）请选择具体原因（必选）包含不实信息涉及个人隐私请选择具体原因（必选）侮辱谩骂诽谤请选择具体原因（必选）搬家样式博文样式补充说明（选填）取消确定.imgViewDom{display:none;position:fixed;top:0;left:0;height:100%;width:100%;z-index:99999999;background:rgba(255,255,255,0.8);overflow:auto;display:-webkit-box;-webkit-box-align:center;-webkit-box-pack:center;display:-moz-box;-moz-box-align:center;-moz-box-pack:center;display:-o-box;-o-box-align:center;-o-box-pack:center;display:-ms-box;-ms-box-align:center;-ms-box-pack:center;display:box;box-align:center;box-pack:center;}.imgViewDomimg{cursor:zoom-out;}新手引导客服举报返回顶部functionshowConfig(){$(window).scrollTop(0);$('.white_content').fadeIn(500);$('body').css('overflow','hidden');$('body').css('filter','blur(3px)');$('body').css('pointer-events','none')}functionsaveAndReload(){$('#configContent').fadeOut(200);setTimeout(function(){location.reload();},200)}"

func TestWord(t *testing.T) {
	seg, _ := gse.New()
	seg.LoadDictEmbed()
	seg.LoadStopEmbed()
	search := seg.CutSearch(testTxt, true)
	for _, s := range search {
		if len(s) > 4 {
			println(s)
		}
	}
}
func TestParseNewWords(t *testing.T) {
	server.WordParseNewWords(testTxt)
}