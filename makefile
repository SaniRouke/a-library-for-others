init:
	@git config --global user.name 'srouke'
	@git config --global user.email 'xpyctam4@yandex.ru'
push:
	@git add .
	@git commit -m "{$1}"
	@git push