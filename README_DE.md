# Einführung

[EN](README.md)


## Der Grund der Entwicklung

Wegen der ehemaligen Arbeit habe ich ein IM-Werkzeug **Slack** verwendet. Es unterstützt bei Integration mit einigen Ereignissen von Gitlab und schickt Nachrichten in der Channels. Die Funktion ist gut, wenn man innerhalb eines Teams miteinander arbeitet.  
Man sollte normalerweise einige **Gruppen** haben, damit man bei der Tätigkeit miteinander über jeweils Projekt kommuniziert, wenn man beruflich WeChatWork benutzt.  
Der Begriff **Gruppe** ist ähnlich wie der Begriff **Channel** bei Slack. Ein **Channel** kann als Projektteam verwendet werden, in dem alle Mitglieder ein Projekt oder mehrere fokussieren müssen.  
Deswegen habe ich mich an Roboter auf WeChatWork erinnert. Ein Roboter benachrichtigt Informationen von den gewählten Repositorien in einer Gruppe, in der es alle Menschen da sind, wer sich mit diesen Projekten beschäftigt. Jeder kann natürlich mehrere **Channels** eintreten, um die neuesten Nachrichten von verschiedenen Projekte zu erhalten. 

## Funktionen

1. Unterstützung bei einigen Events von Gitlab
* Push Event
* Issue Event
* Merge Request Event
2. Konfiguration für den ausgewählten Roboter durch Path Variable


# Anweisung
1. Bitte Administrator, um ein neuer Roboter zu erstellen und erhalte die Identität davon.
2. Setzt „Webhook“ unter Settings -> Integrations des Git-Repositoriums ein.
```
URL: http://host:port/webhooks/gitlab/${robotId}
Secret Token: chatten mit mir
Trigger: Kreuz die entsprechenden „Events“ je nach den Bedürfnissen
```
3. Los geht's!
