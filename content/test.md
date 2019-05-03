+++
Title: Test
Category: 1
Hero: /static/img/hero-14.png
Publish: <nil>
+++

Here are some headers with text to keep you going {:truncate}

# Wow this is important
By Thursday, you have demonstrated that you are the strongest, smartest, and most passionate engineer on your team. When you get to your desk, rip your standard-issue keyboard out and slide it off your desk. Plug in your own mechanical keyboard with Cherry MX Green or Blue switches. Your typing will now be a constant audible reminder to your teammates that you aren’t fucking around.{:longtruncate}
Log into your team’s bug-tracking software and look around for some low-hanging fruit to fix. Spend the rest of the day working and fixing them, but don’t send any code reviews out.

![image-center](https://s3.amazonaws.com/img.rawrrawr.com/gallery/fbd0e6258a5f40832e2c2e4e7fca7700d20755a3.jpeg "This should be the title text underneat the image! Its a kitty!")

As you are speaking, lock eye-contact with the person to your left, and do not break until they look away. At this point, lock eye-contact with the person to their left, and do the same. By the time you are finished speaking, you will have gone full circle and sent a clear alpha signal to each of them. Perform this ritual at every stand-up.
If someone tells you to take it offline while you are speaking, brush it off. Inform them that what are you saying is the most important thing that has ever been spoken at this company. Then, pull the red pen out of your pocket and scribble their name down.

Anyway, **bold text**  and some _underlined text_, and maybe both? _**wow i'm upset about something**_

-----

## Second Paragraph
As you are working, make sure to look extra pissed off at all times, like you can’t believe how your teammates managed to write such crappy code. Type more and more furiously as the day progresses. Do not eat lunch. Do not take breaks. Only leave your desk when you are the last man in the office.

### Third paragraph
At around 10 PM, remote in from home and send your first CR out to your team. In an hour, send another CR out. You should have fixed enough bugs during the day to do this for the next four to five hours. You are one hard-working, dedicated, bad-ass motherfucker.

#### Fourth Paragraph
Wednesday is time to institute your technical sovereignty. Be the first one at your daily stand-up, and run it like the Scrum master you know you are. If someone talks for more than ten seconds, immediately interrupt them and tell them to take it offline. Bring a notepad and take notes with a black pen. Keep a second red pen in your pocket.

* Backwards talking!!
* A thing?
* Wow why is this
* list of many things

> When it’s your turn to speak, go on a long rant about all the horrible design patterns you’ve discovered in the code on Tuesday. Announce that you have already started designing the architecture of the inevitable re-write. Name-drop as many of the latest software frameworks and technologies as possible throughout your rant. Use words like big data, cloud, and scalability. Mention test-driven development at least three to four times.

Oh right, here is a [Link to google](http://www.google.com), this needs to be styled a bit, don't you think?

```cpp
if (material.constantBuffer->flags & FLAGS_BUFFER_TRANSIENT)
{
	for (size_t i = 0; i < TRANSIENT_CONSTANT_BUFFER_COUNT; ++i)
	{
		if (material.constantBuffer->size < TRANSIENT_CONSTANT_BUFFER_SIZES[i])
		{
			cb = transientConstantBuffers[i];
			break;
		}
	}

	assert(cb);
	D3D11_MAPPED_SUBRESOURCE subresource;

	HRESULT hr = context->Map(cb, 0, D3D11_MAP_WRITE_DISCARD, 0, &subresource);
	if (FAILED(hr))
	{
		throw Exceptions::DirectXError(hr, "Context::Map");
	}

	material.constantBuffer->fillTransient(subresource.pData, transform.data(), 16);
	context->Unmap(cb, 0);
}
```

Consider the matrix $$A = [[a, b],[c,d]]$$, and vector $$ vec v = (1, 2)$$. Then $$ vecvA = (a + 2c, b + 2d)$$

As you are speaking, lock eye-contact with the person to your left, and do not break until they look away. At this point, lock eye-contact with the person to their left, and do the same. By the time you are finished speaking, you will have gone full circle and sent a clear alpha signal to each of them. Perform this ritual at every stand-up.
If someone tells you to take it offline while you are speaking, brush it off. Inform them that what are you saying is the most important thing that has ever been spoken at this company. Then, pull the red pen out of your pocket and scribble their name down.

Element | Bytes | Description
--------|-------|-------------------------------------
n       | 4     | Number of characters to follow.
str     | n     | Characters.
off     | 2     | Offset of first value
t       | 1     | A selector byte.

Anyway, here is some "double quoted text", some 'single quoted text', and one dash - two dashes -- and three dashes --- finally.