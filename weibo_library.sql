-- MySQL dump 10.13  Distrib 5.5.27, for osx10.6 (i386)
--
-- Host: localhost    Database: weibo_library
-- ------------------------------------------------------
-- Server version	5.5.27-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `books` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(50) NOT NULL,
  `name` char(100) NOT NULL,
  `author` char(50) DEFAULT NULL,
  `info` text,
  `img` char(200) DEFAULT NULL,
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `count` int(20) unsigned DEFAULT '1',
  `status` tinyint(1) unsigned DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `author_index` (`author`),
  KEY `user_index` (`username`),
  KEY `book_index` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8 COMMENT='books';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (29,'daodao','Big Go','ken',NULL,NULL,'2013-02-03 11:21:25',1,1),(40,'5','a','1','1','','2013-02-05 02:25:17',0,0),(41,'5','aaa','1','1','','2013-02-05 02:26:19',0,0),(42,'5','ttt','1','2','','2013-02-05 02:31:07',0,0),(43,'5','r','','','','2013-02-05 02:32:17',0,0),(44,'5','sss','','','','2013-02-05 02:34:51',0,0),(45,'5','uuu','','','','2013-02-05 02:34:55',0,0),(46,'5','as','','','','2013-02-05 02:36:09',0,0),(47,'5','umit','','','','2013-02-05 02:36:14',0,1),(48,'5','sfff','','','','2013-02-05 03:15:56',0,1),(49,'5','2222','222','','','2013-02-11 09:03:45',0,1),(50,'5','666','daodao','','','2013-02-11 10:58:28',0,1),(51,'5','ggg','www','','','2013-02-11 11:07:05',0,1),(52,'5','sssw','','','','2013-02-11 11:07:24',0,1),(53,'5','五千年','司马迁','','','2013-02-11 11:33:10',0,1),(54,'5','时间简史','道道','','','2013-02-11 11:34:34',0,1),(55,'daodao','big table','cutting','','','2013-02-12 01:26:41',0,1),(56,'daodao','ssd','ssd','','','2013-02-12 01:27:44',0,1),(57,'1','book1','a','','','2013-02-12 01:43:33',0,1),(58,'c','object c教程','刀道','好书一本','','2013-02-16 06:10:13',0,1),(59,'guanpu','hadoop实战','','hadoop入门','','2013-02-18 00:55:25',0,1),(60,'guanpu','数学之美','吴军','了解数学的奥妙','','2013-02-18 00:55:49',0,0),(61,'test','abc','','','','2013-02-18 01:25:18',0,2),(62,'test2','abc','','','','2013-02-18 01:33:39',0,1),(63,'shijin1','test','a','a','','2013-02-18 01:42:55',0,1),(64,'wenhong1','[head frist] Python','','','','2013-02-18 01:47:30',0,0),(65,'wenhong1','TCP/IP 详解 卷1','','','','2013-02-18 01:50:01',0,0),(66,'wenhong1','TCP/IP 详解 卷2 ：实现','','','','2013-02-18 01:50:14',0,0),(67,'wenhong1','TCP/IP 详解 卷3','','','','2013-02-18 01:50:34',0,0),(68,'tanzheng','如何解题','','','','2013-02-18 01:55:35',0,0),(69,'tanzheng','失控','凯文.凯利','','','2013-02-18 01:57:14',0,0),(70,'junwei','Linux Shell脚本攻略','','linux shell非常不错的一本书','','2013-02-18 02:02:34',0,0),(71,'junwei','hadoop实战','','','','2013-02-18 02:02:50',0,0),(72,'junwei','java编程思想','','','','2013-02-18 02:03:01',0,0),(73,'junwei','古炉','贾平凹','推荐','','2013-02-18 02:03:18',0,0),(74,'junwei','高老庄','贾平凹','','','2013-02-18 02:03:30',0,0),(75,'junwei','从文自传','沈从文','强烈推荐','','2013-02-18 02:03:55',0,0),(76,'junwei','营养圣经','','看过，不错，值得一看，了解人体所需要的营养元素','','2013-02-18 02:04:30',0,0),(77,'guanpu','Go语言编程','许式伟','读书圈网站的指导资料之一  入门Go语言，不过最好的地方还是自己下载 跑起来 godoc','','2013-02-18 03:33:14',0,0),(78,'yangwm','程序员修炼之道--从小工到专家','Andrew Hunt','','','2013-02-19 11:04:34',0,0),(79,'yangwm','Debug Hacks中文版','吉冈','深入调试的技术与工具','','2013-02-19 11:06:38',0,0),(80,'guanpu','一万小时天才理论','丹尼尔.科伊尔','把一个事情当作终生的事业。把编程当作终生的事业，那么就不会那么浅视，也会更有计划，也会爱惜自己的身体。','','2013-02-20 00:32:12',0,0),(81,'guanpu','树上的男爵','卡尔维诺','诺贝尔奖获得者卡尔维诺，意大利风情，树上生活的一辈子，另一种角度看人生','','2013-02-22 01:03:04',0,0),(82,'guanpu','如何阅读一本书','莫提默·J. 艾德勒 / 查尔斯·范多伦','《如何阅读一本书》初版于1940年，1972年大幅增订改写为新版。不懂阅读的人，初探阅读的人，读这本书可以少走冤枉路。对阅读有所体会的人，读这本书可以有更深的印证和领悟。','','2013-02-22 01:04:32',0,0),(83,'guanpu','七周七语言','Bruce A.Tate','2011年Jolt大奖图书 带你入门 带你开阔视野','','2013-02-25 01:55:26',0,0);
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `borrow_info`
--

DROP TABLE IF EXISTS `borrow_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `borrow_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `book_id` bigint(20) unsigned NOT NULL,
  `owner_id` char(200) NOT NULL,
  `borrower_id` char(50) NOT NULL,
  `note` text NOT NULL,
  `start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `is_back` tinyint(1) unsigned DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `book_index` (`book_id`),
  KEY `owner_index` (`owner_id`),
  KEY `borrow_index` (`borrower_id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8 COMMENT='borrow_info';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `borrow_info`
--

LOCK TABLES `borrow_info` WRITE;
/*!40000 ALTER TABLE `borrow_info` DISABLE KEYS */;
INSERT INTO `borrow_info` VALUES (32,60,'43','47','1 数学在上学的时候觉得没什么用，工作后用到的也很少，即使算法也都是一些基本的。但是看完数学之美才发现数学原来这么有用。就像文中所说：数学是上帝描写自然的语言。纯数学使我们能够发现概念和联系这些概念的规律，这些概念和规律给了我们理解自然现象的钥匙。  2 统计语言模型，阿兰•图灵提出计算的机器和智能。先有＂鸟飞派＂，用基于规则的方法来学习人类，但发现走不通。改用基于统计的处理方法，IBM科学家卡茨发明了卡茨退避法，来计算概率，根据大数定理只要统计量足够，相对频度就等于概率。看来IBM还是真是强大。\r\n     卡茨退避法（Katz backoff），对于频率超过一定阈值的词，它们的概率估计就是它们在语料库中的相对频度，对于频率小于这个阈值的词，它们的概率估计就小于他们的相对频度，出现次数越少，频率下调越多。对于未看见的词，也给予一个比较小的概率（即下调得到的频率总和），这样所有词的概率估计都平滑了。这就是卡茨退避法（Katz backoff）。\r\n4 中文分词\r\n     利用统计语言模型，只能解决大多数问题，无法消除二义性。需要在词的颗粒度和层次划分上做文章。\r\n5 隐含马尔可夫模型是机器学习的主要工具，需要结合训练算法鲍姆-韦尔奇算法和解码算法维特比算法。\r\n     隐含马尔科夫模型是马尔科夫链的扩展，任意时刻t的状态st是不可见的，所以观察者没法通过观察到一个状态序列s1,s2,s3,…,sT来推测转移概率等参数。但是隐马尔科夫模型在每个时刻t会输出一个符号ot，而且ot和st相关且仅和ot相关。这个被称为独立输出假设。其中隐含的状态s1,s2,s3,…是一个典型的马尔科夫链。\r\n     \r\n     鲍姆-韦尔奇算法（Baum-Welch Algorithm），首先找到一组能够产生输出序列O的模型参数，这个初始模型成为Mtheta0，需要在此基础上找到一个更好的模型，假定不但可以算出这个模型产生O的概率P(O|Mtheta0)，而且能够找到这个模型产生O的所有可能的路径以及这些路径的概率。并算出一组新的模型参数theta1，从Mtheta0到Mtheta1的过程称为一次迭代。接下来从Mtheta1出发寻找更好的模型Mtheta2，并一直找下去，直到模型的质量没有明显提高为止。这样一直估计（Expectation）新的模型参数，使得输出的概率达到最大化（Maximization）的过程被称为期望值最大化（Expectation-Maximization）简称EM过程。EM过程能保证一定能收敛到一个局部最优点，但不能保证找到全局最优点。因此，在一些自然语言处理的应用中，这种无监督的鲍姆-韦尔奇算法训练处的模型比有监督的训练得到的模型效果略差。\r\n     维特比算法（Viterbi Algoritm）是一个特殊但应用最广的动态规划算法，利用动态规划，可以解决任何一个图中的最短路径问题。它之所以重要，是因为凡是使用隐含马尔科夫模型描述的问题都可以用它来解码。1.从点S出发，对于第一个状态x1的各个节点，不妨假定有n1个，计算出S到他们的距离d(S,x1i)，其中x1i代表任意状态1的节点。因为只有一步，所以这些距离都是S到他们各自的最短距离。2.对于第二个状态x2的所有节点，要计算出从S到他们的最短距离。d(S,x2i)=min_I=1,n1_d(S,x1j)+d(x1j,x2i)，由于j有n1种可能性，需要一一计算，然后找到最小值。这样对于第二个状态的每个节点，需要n1次乘法计算。假定这个状态有n2个节点，把S这些节点的距离都算一遍，就有O(n1*n2)次运算。3.按照上述方法从第二个状态走到第三个状态一直走到最后一个状态，这样就得到整个网络从头到尾的最短路径。\r\n6 信息的度量和作用\r\n      信息的作用在于消除不确定性。\r\n      熵，是对信息的量化度量。信息熵的定义为H(X)=-SumP(x)logP(x)，变量的不确定性越大，熵也越大。\r\n      互信息，对两个随机事件相关性的量化度量，即随机事件X的不确定性或者说熵H(X)，在知道随机事件Y条件下的不确定性，或者说条件熵H(X|Y)之间的差异，即I(X;Y)=H(X)-H(X|Y)。所谓两个事件相关性的量化度量，即在了解了其中一个Y的前提下，对消除另一个X不确定性所提供的信息量。\r\n      相对熵（Kullback-Leibler Divergence）也叫交叉熵，对两个完全相同的函数，他们的相对熵为零；相对熵越大，两个函数差异越大，反之，相对熵越小，两个函数差异越小；对于概率分布或者概率密度函数，如果取值均大于零，相对熵可以度量两个随机分布的差异性。\r\n7 简单之美\r\n     布尔代数只有0和1两个数字,但所有数学和逻辑运算都可以转换成布尔运算。 真理在形式上从来是简单的，而不是复杂和含混的。\r\n     采用布尔代数，搜索引擎不过是一张大表，表的每一行对应一个关键字，而每一个关键字后面跟着一组数字，是包含该关键词的文献序号。但当索引变的非常大的时候，这些索引需要通过分布式的方式存储到不同的服务器上。\r\n8 图论和网络爬虫\r\n     网络爬虫（Web Crawlers），图论的遍历算法和搜索引擎的关系。互联网虽然复杂，但是说穿了其实就是一张大图……可以把每一个网页当做一个节点，把那些超链接当做连接网页的弧。有了超链接，可以从任何一个网页出发，用图的遍历算法，自动访问到每一个网页并且把他们存储起来。完成这个功能的程序叫网络爬虫。     哥尼斯堡七桥，如果一个图能从一个顶点出发，每条边不重复的遍历一遍回到这个顶点，那么每一个顶点的度必须为偶数。     构建网络爬虫的工程要点：1.用BFS（广度优先搜索）还是DFS（深度优先搜索），一般是先下载完一个网站，再进入下一个网站，即BFS的成分多一些。2.页面的分析和URL的提取，如果有些网页明明存在，但搜索引擎并没有收录，可能的原因之一是网络爬虫中的解析程序没能成功解析网页中不规范的脚本程序。3.记录哪些网页已经下载过的URL表，可以用哈希表。最终，好的方法一般都采用了这样两个技术：首先明确每台下载服务器的分工，也就是在调度时，一看到某个URL就知道要交给哪台服务器去下载，这样就避免了很多服务器对同一个URL做出是否需要下载的判断。然后，在明确分工的基础上，判断URL是否下载就可以批处理了，比如每次向哈希表（一组独立的服务器）发送一大批询问，或者每次更新一大批哈希表的内容，这样通信的次数就大大减少了。\r\n9 PageRank\r\n     PageRank衡量网页质量的核心思想，在互联网上，如果一个网页被很多其他网页所链接，说明它受到普遍的承认和信赖，那么它的排名就高。同时，对于来自不同网页的链接区别对待，因为网页排名高的那些网页的链接更可靠，于是要给这些链接比较大的权重。\r\n      采用了算法，TF-IDF(Term Frequency / Inverse Document Frequency) ，关键词频率-逆文本频率值，其中，TF为某个网页上出现关键词的频率，IDF为假定一个关键词w在Dw个网页中出现过，那么Dw越大，w的权重越小，反之亦然，公式为log(D/Dw)。1.一个词预测主题的能力越强，权重越大，反之，权重越小。2.停止词的权重为零。\r\n10 有限状态机和动态规划\r\n     有限状态机（finite-state machine, FSM），又称有限状态自动机，简称状态机，是表示有限个状态以及在这些状态之间的转移和动作等行为的数学模型。     \r\n\r\n     动态规划（Dynamic Programming）的原理，将一个寻找全程最优的问题分解成一个个寻找局部最优的小问题。\r\n11 新闻分类\r\n     在实际的分类中，可以先进行奇异值分解（得到分类结果略显粗糙但能较快得到结果），在粗分类结果的基础上，利用计算向量余弦的方法（对范围内的分类做两两计算），在粗分类结果的基础上，进行几次迭代，得到比较精确的结果。\r\n\r\n12 信息指纹\r\n     信息指纹。如果能够找到一种函数，将5000亿网址随即地映射到128位二进制，也就是16字节的整数空间，就称这16字节的随机数做该网址的信息指纹。信息指纹可以理解为将一段信息映射到一个多维二进制空间中的一个点，只要这个随即函数做的好，那么不同信息对应的点不会重合，因此这个二进制的数字就变成了原来信息所具有的独一无二的指纹。\r\n     判断两个集合是否相同，最笨的方法是这个集合中的元素一一比较，复杂度O(squareN)，稍好的是将元素排序后顺序比较，复杂度O(NlogN)，最完美的方法是计算这两个集合的指纹，然后直接进行比较，计算复杂度O(N)。\r\n     YouTube 反盗版，利用关键帧提取和特征提取组成信息指纹进行视频匹配，效果明显。\r\n\r\n13 最大熵\r\n     最大熵模型（Maximum Entropy）的原理就是保留全部的不确定性，将风险降到最小。最大熵原理指出，需要对一个随机事件的概率分布进行预测时，我们的预测应当满足全部已知的条件，而对未知的情况不要做任何主观假设。在这种情况下，概率分布最均匀，预测的风险最小。I.Csiszar证明，对任何一组不自相矛盾的信息，这个最大熵模型不仅存在，而且是唯一的，此外，他们都有同一个非常简单的形式-指数函数。\r\n     \r\n     当年最早改进最大熵模型算法的达拉皮垂兄弟这些年难道没有做任何事吗？他们在九十年代初贾里尼克离开 IBM 后，也退出了学术界，而到在金融界大显身手。他们两人和很多 IBM 语音识别的同事一同到了一家当时还不大，但现在是世界上最成功对冲基金(hedge fund)公司----文艺复兴技术公司 (Renaissance Technologies)。我们知道，决定股票涨落的因素可能有几十甚至上百种，而最大熵方法恰恰能找到一个同时满足成千上万种不同条件的模型。达拉皮 垂兄弟等科学家在那里，用于最大熵模型和其他一些先进的数学工具对股票预测，获得了巨大的成功。从该基金 1988 年创立至今，它的净回报率高达平均每年 34%。也就是说，如果 1988 年你在该基金投入一块钱，今天你能得到 200 块钱。这个业绩，远远超过股神巴菲特的旗舰公司伯克夏哈撒韦（Berkshire Hathaway)。同期，伯克夏哈撒韦的总回报是 16 倍。\r\n14 输入法的发展\r\n     以词为单位做单位统计信息熵，结合上下文。拼音转汉字使用动态规划的数学模型。通过用户输入的内容训练语言模型。\r\n     可见要把用户当成小白，不能要求用户做任何额外的东西，才能吸引用户的使用。\r\n15 布隆过滤器\r\n     布隆过滤器实际上是一个很长的二进制向量和一系列随机映射的函数。主要用于判断一个元素是否包含在一个集合中。\r\n16 贝叶斯网络\r\n     贝叶斯网络从数学的层面讲是一个加权的有向图，是马尔科夫链的扩展，而从知识论的层面看，贝叶斯网络克服了马尔科夫那种机械的线性的约束，它可以把任何有关联的事件统一到它的框架下面。在网络中，假定马尔科夫假设成立，即每一个状态只与和它直接相连的状态有关，而和他间接相连的状态没有直接关系，那么它就是贝叶斯网络。在网络中每个节点概率的计算，都可以用贝叶斯公式来进行，贝叶斯网络也因此得名。由于网络的每个弧都有一个可信度，贝叶斯网络也被称作信念网络（Belief Networks）。\r\n     用于基于统计的模型分析文本，从中抽取概念，分析主题。\r\n17 条件随机场\r\n     条件随机场是计算联合概率分布的有效模型。在一个隐含马尔科夫模型中，以x1,x2,...,xn表示观测值序列，以y1,y2,...,yn表示隐含的状态序列，那么xi只取决于产生它们的状态yi,和前后的状态yi-1和yi+1都无关。显然很多应用里观察值xi可能和前后的状态都有关，如果把xi和yi-1,yi,yi+1都考虑进来，这样的模型就是条件随机场。它是一种特殊的概率图模型（Probablistic Graph Model），它的特殊性在于，变量之间要遵守马尔科夫假设，即每个状态的转移概率只取决于相邻的状态，这一点和另一种概率图模型贝叶斯网络相同，它们的不同之处在于条件随机场是无向图，而贝叶斯网络是有向图。\r\n18 期望最大化算法（Expectation Maximization Algorithm），根据现有的模型，计算各个观测数据输入到模型中的计算结果，这个过程称为期望值计算过程（Expectation），或E过程；接下来，重新计算模型参数，以最大化期望值，这个过程称为最大化的过程（Maximization），或M过程。这一类算法都称为EM算法，比如隐含马尔科夫模型的训练方法Baum-Welch算法，以及最大熵模型的训练方法GIS算法。\r\n     EM算法只需要有一些训练数据，定义一个最大化函数，剩下就是计数机的事情，经过若干次迭代就得到结果。\r\n19  逻辑回归模型（Logistic Regression）是将一个事件出现的概率适应到一条逻辑曲线（Logistic Curve）上。典型的逻辑回归函数：f(z)=e`z/e`z+1=1/1+e`-z。逻辑曲线是一条S型曲线，其特点是开始变化快，逐渐减慢，最后饱和。逻辑自回归的好处是它的变量范围从负无穷到正无穷，而值域范围限制在0-1之间。因为值域的范围在0-1之间，这样逻辑回归函数就可以和一个概率分别联系起来了。因为自变量范围在负无穷到正无穷之间，它就可以把信号组合起来，不论组合成多大或者多小的值，最后依然能得到一个概率分布。\r\n     一个广告系统中，有没有好的点击率预估机制决定是否能成倍提高单位搜索的收入。\r\n20 云计算\r\n     从分治算法到MapReduce。\r\n     Google针对云计算给出的解决工具是MapReduce，其根本原理就是计算机算法上常见的分治算法（Divide-and-Conquer）。将一个大任务拆分成小的子任务，并完成子任务的计算，这个过程叫Map，将中间结果合并成最终结果，这个过程叫Reduce。\r\n     \r\n','2013-02-18 01:40:56','0000-00-00 00:00:00',1),(33,59,'43','48','','2013-02-18 01:46:23','0000-00-00 00:00:00',0),(34,63,'47','48','','2013-02-18 01:47:54','0000-00-00 00:00:00',1),(35,67,'48','49','test','2013-02-18 01:50:41','0000-00-00 00:00:00',1),(36,65,'48','50','www','2013-02-18 01:51:32','0000-00-00 00:00:00',1),(37,63,'47','55','','2013-02-19 11:02:45','0000-00-00 00:00:00',0),(38,83,'43','57','','2013-02-25 01:57:25','0000-00-00 00:00:00',1);
/*!40000 ALTER TABLE `borrow_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pwds`
--

DROP TABLE IF EXISTS `pwds`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pwds` (
  `username` char(50) NOT NULL,
  `pwd` char(50) NOT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='passwords';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pwds`
--

LOCK TABLES `pwds` WRITE;
/*!40000 ALTER TABLE `pwds` DISABLE KEYS */;
/*!40000 ALTER TABLE `pwds` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(50) NOT NULL,
  `name` char(50) DEFAULT NULL,
  `own_book_count` int(10) NOT NULL DEFAULT '0',
  `borrow_book_count` int(10) NOT NULL DEFAULT '0',
  `allow_book_count` int(10) NOT NULL DEFAULT '0',
  `allow_borrow_days` tinyint(10) NOT NULL DEFAULT '30',
  `user_level` tinyint(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`username`),
  KEY `name` (`name`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8 COMMENT='user';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (45,'','',0,0,0,30,0),(36,'2','2',0,0,0,30,0),(40,'a','a',0,0,0,30,0),(37,'aa','aa',0,0,0,30,0),(38,'aaa','aaa',0,0,0,30,0),(33,'abc','abc',0,0,0,30,0),(41,'b','b',0,0,0,30,0),(39,'bbb','bbb',0,0,0,30,0),(43,'guanpu','guanpu',7,2,0,30,0),(57,'haitao5','haitao5',0,1,0,30,0),(52,'jiankun','jiankun',0,0,0,30,0),(49,'jintao1','jintao1',0,1,0,30,0),(53,'junwei','junwei',7,0,0,30,0),(47,'shijin1','shijin1',1,1,0,30,0),(51,'tanzheng','tanzheng',2,0,0,30,0),(56,'weijia7','weijia7',0,0,0,30,0),(48,'wenhong1','wenhong1',4,2,0,30,0),(30,'www','好人',0,0,0,30,0),(54,'xiaobing2@staff.sina.com.cn','xiaobing2@staff.sina.com.cn',0,0,0,30,0),(50,'yangjing1@staff.sina.com.cn','yangjing1@staff.sina.com.cn',0,1,0,30,0),(55,'yangwm','yangwm',2,1,0,30,0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_pwd`
--

DROP TABLE IF EXISTS `user_pwd`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_pwd` (
  `username` char(50) NOT NULL,
  `pwd` char(50) NOT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='passwords';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_pwd`
--

LOCK TABLES `user_pwd` WRITE;
/*!40000 ALTER TABLE `user_pwd` DISABLE KEYS */;
INSERT INTO `user_pwd` VALUES ('a','%x[144 1 80 152 60 210 79 176 214 150 63 125 40 22'),('b','900150983cd24fb0d6963f7d28e17f72'),('bbb','32'),('c','c4ca4238a0b923820dcc509a6f75849b'),('junwei','b86625a4e9a1bc96f574e6dc9007de29'),('test','c4ca4238a0b923820dcc509a6f75849b'),('test2','c81e728d9d4c2f636f067f89cc14862c'),('wenhong1','f239bc12ce90dbcb5eda5f82a6d0d64a'),('yangwm','e10adc3949ba59abbe56e057f20f883e');
/*!40000 ALTER TABLE `user_pwd` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2013-03-05 21:03:34
