// cada servicio tambien puede tener su middleware

Como usar e abort en los midlewares

// ? la porpose se verifica en el handler que lo requiere

```
/* solo un valor pasado
		ctx := context.WithValue(r.Context(), UserIDKey, c.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
		*/
		/* este parece sin pasar valores
		_, err =
		next.ServeHTTP(w,r)
		*/
```

```
ctx := context.WithValue(r.Context(), UserIDKey, c.Subject)
		ctx = context.WithValue(ctx, TokenTypeIDKey, c.TokenType)
		//ctx = context.WithValue(ctx,AnotherIDKey, c.Another)
		next.ServeHTTP(w, r.WithContext(ctx))
```
