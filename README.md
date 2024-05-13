<h2 align="center">Gitleaks pre-commit hook</h2>

### Використання pre-commit хука з gitleaks з локальною інсталяцією gitleaks (curl)

### Встановлення та використання

1. Увімкніть pre-commit хук:

```bash
$ git config hooks.pre-commit.enable true
```

2. Pre-commit хук знаходиться у папці hooks. Перенести скрипт до хуків git, або виконати команду:

```bash
$ cp hooks/pre-commit .git/hooks/pre-commit
```

4. Внести зміни до коду і перевірити використання. Може буди необхідним root права для встановлення gitleaks при першому використанні

```bash
$ git add .
$ git commit -m "test"
```

1. Вимкнення pre-commit хука:

```bash
 $ git config hooks.pre-commit.enable false
```
