<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Directory structure
<!--endTocHeader-->

```
.
├── Dockerfile
├── README.md
├── _docs
├── alembic
│   ├── README
│   ├── env.py
│   ├── script.py.mako
│   └── versions
│       └── 7dda1641a129_20220611183010_init.py
├── alembic.ini
├── avro
├── configs
│   ├── __init__.py
│   ├── appFactory.py
│   ├── auth.py
│   ├── cors.py
│   ├── db.py
│   ├── dir.py
│   ├── error.py
│   ├── featureFlag.py
│   ├── kafka.py
│   ├── menuServiceFactory.py
│   ├── messagebus.py
│   ├── messagebusFactory.py
│   ├── pageTemplateFactory.py
│   ├── port.py
│   ├── rmq.py
│   ├── rpc.py
│   ├── rpcFactory.py
│   ├── ui.py
│   └── url.py
├── helpers
│   ├── __init__.py
│   └── transport
├── main.py
├── migrate.sh
├── modules
│   ├── __init__.py
│   ├── auth
│   │   ├── __init__.py
│   │   ├── auth
│   │   ├── event.py
│   │   ├── role
│   │   ├── route.py
│   │   ├── rpc.py
│   │   ├── session
│   │   ├── token
│   │   ├── user
│   │   └── userSeeder
│   └── ui
│       ├── __init__.py
│       ├── menu
│       └── page
├── pages
│   ├── auth
│   │   └── crud
│   ├── default-partials
│   │   ├── include-css.html
│   │   ├── include-js.html
│   │   ├── meta.html
│   │   └── navigation.html
│   ├── default_crud.html
│   ├── default_error.html
│   ├── default_login.html
│   ├── default_logout.html
│   └── default_page.html
├── public
│   ├── css
│   ├── favicon
│   ├── js
│   │   ├── app.js
│   │   └── vue-sfc-loader-options.js
│   └── vue
│       └── auth
├── repos
│   ├── __init__.py
│   └── base.py
├── requirements.txt
├── schemas
│   ├── __init__.py
│   ├── authType.py
│   ├── menu.py
│   ├── menuContext.py
│   ├── role.py
│   ├── test_role.py
│   ├── test_user.py
│   └── user.py
├── start.sh
├── structure.txt
└── template.env

```

TODO: Write about `Directory structure`

<!--startTocSubTopic-->
<!--endTocSubTopic-->