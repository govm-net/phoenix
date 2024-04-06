# phoenix

**phoenix** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```bash
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project.

For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```bash
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```bash
curl https://get.ignite.com/govm-net/phoenix@latest! | sudo bash
```

`govm-net/phoenix` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

## virtual block

用于不同分片之间的相互锁定和同步，这是虚拟区块，和cosmos的block不会是一一映射的。
cosmos的block可能是1秒一个，而vBlock可能是10秒一个。vBlock不需要太频繁，主要是用于不同分片之间的交互

1. LastHeader:cosmos block info
2. preBlock:
3. parentBlock
4. leftChildBlock
5. rightChildBlock
6. index
7. timestramp
8. vdf

## shard

允许发起投票，创建分片

1. 要求分片id为1
2. 扣除手续费
3. 将手续费作为新分片矿工的质押资产
4. 父分片的矿工将作为新分片的初始矿工
5. 要求父分片已经创建超过7天

通过app.MsgServiceRouter()，可以找到message的handler，就可以调用处理message。
keeper保存MsgServiceRouter，参考gov的router。MsgSubmitProposalCmd

应该再增加两个map（toParent,toChild），用于判断是否支持对应的message，避免滥用导致异常。Key为message.Type。

## todo

### block

1. 时间戳：禁止区块间隔时间太长
2. 避免长时间不生成区块，如果出现，应该用空区块代替
   1. 空区块用固定的时间戳，不包含任何交易
3. 限制投票支持的消息类型，避免分片恶意mint token等行为
4. 父链包含子链的投票权
