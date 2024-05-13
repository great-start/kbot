<h2 align="center">Gitleaks pre-commit hook</h2>

### Використання pre-commit хука з gitleaks з локальною інсталяцією gitleaks (curl)

### Встановлення та використання

1. Увімкніть pre-commit:

```bash
$ git config hooks.pre-commit.enable true
```

2. Перенести скрипт до хуків git

```bash
$ cp hooks/pre-commit .git/hooks/pre-commit
```

4. Внести зміни до коду і перевірити використання

```bash
$ git add .
$ git commit -m "test"
```

1. Вимкнення pre-commit хука:

```bash
 $ git config hooks.pre-commit.enable false
```
