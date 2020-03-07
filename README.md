# abstraction-of-transaction

UnitOfWork Pattern を用いての Transaction の抽象化

[repository-with-tx](https://github.com/trewanek/repository-with-tx)

との違いは、

- Rollback / Commit が UseCase 内に明示的に現れている
- Infrastructure 依存が UseCase から消え、UnitOfWork に移譲されている点

UnitOfWork を New した段階で Transaction は開始される。

## 疑問点

- UnitOfWork の実装はどの層におくべきか(Transaction の管理主体が UseCase なので Application 層だと思っている)
- UnitOfWork が Infrastructure に依存している(Application 層に置く場合は、依存関係的に正しくなさそう)

## 参考

- [ボトムアップドメイン駆動設計 後編](https://nrslib.com/bottomup-ddd-2/)
- [ドメイン駆動設計入門 ボトムアップでわかる！ドメイン駆動設計の基本](https://www.amazon.co.jp/dp/B082WXZVPC)
