<!--startTocHeader-->
[🏠](../README.md) > [Adding a new module](README.md)
# Adding a page
<!--endTocHeader-->


You can use [Zaruba](https://github.com/state-alchemists/zaruba) to add a new page, or you can write a page from scratch.

# Using Zaruba

To create a new module using Zaruba, you can invoke the following code:

```bash
zaruba please addFastAppPage \
    appDirectory=ztplAppDirectory \
    appModuleName=yourModuleName \
    appHttpMethod=get \
    appUrl=/your-end-point
# or:
# zaruba please addFastAppPage -i
```

Once created, you can go to `pages/modules/your-end-point.html` and start editing the HTML.

# From scratch

To add a page from scratch, you need to edit `module/<module-name>/route.py`.

```python
# location: module/library/route.py

# rest of the code...

################################################
# -- 👓 User Interface
################################################
def register_library_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    menu_service.add_menu(name='library', title='Library', url='#', auth_type=AuthType.ANYONE)

    # Blog page
    menu_service.add_menu(name='blog:/blog', title='Blog', url='/library/blog', auth_type=AuthType.ANYONE, parent_name='library')
    @app.get('/library/blog', response_class=HTMLResponse)
    async def get_blog(request: Request, context: MenuContext = Depends(menu_service.has_access('blog:/blog'))) -> HTMLResponse:
        '''
        Serve (get) /blog
        '''
        try:
            return page_template.TemplateResponse('default_page.html', context={
                'request': request,
                'context': context,
                'content_path': 'modules/library/blog.html'
            }, status_code=200)
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            return page_template.TemplateResponse('default_error.html', context={
                'request': request,
                'status_code': 500,
                'detail': 'Internal server error'
            }, status_code=500)


    register_book_ui_route(app, mb, rpc, menu_service, page_template)

    print('Register library UI route handler', file=sys.stderr)
```

Then you need to create `pages/modules/library/blog.html`

```html
<p>
    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
</p>
```

<!--startTocSubTopic-->
<!--endTocSubTopic-->