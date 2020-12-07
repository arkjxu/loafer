package loafer

const (
	// INSTALLSUCCESSPAGE - Default Installation Page
	INSTALLSUCCESSPAGE = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
			<style>
				@font-face {
					font-family: 'iconfont';  /* project id 2243819 */
					src: url('//at.alicdn.com/t/font_2243819_i7pd6nsp1wp.eot');
					src: url('//at.alicdn.com/t/font_2243819_i7pd6nsp1wp.eot?#iefix') format('embedded-opentype'),
					url('//at.alicdn.com/t/font_2243819_i7pd6nsp1wp.woff2') format('woff2'),
					url('//at.alicdn.com/t/font_2243819_i7pd6nsp1wp.woff') format('woff'),
					url('//at.alicdn.com/t/font_2243819_i7pd6nsp1wp.ttf') format('truetype'),
					url('//at.alicdn.com/t/font_2243819_i7pd6nsp1wp.svg#iconfont') format('svg');
				}
				.iconfont {
					font-family: 'iconfont';
				}
				html, body, .install-successful {
					height: 100%;
					width: 100%;
					
				}
				.install-successful {
					display: flex;
					justify-content: center;
					align-items: center;
				}
				.success {
					color: #3FDF64;
					font-size: 2rem;
					font-style: normal;
					padding: 0 8px;
				}
				.card {
					box-shadow: 0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23);
					padding: 20px 50px;
				}
			</style>
		</head>
		<body>
			<div class="install-successful">
				<div class="card">
					<div style="display: flex; flex-flow: row nowrap; justify-content: center; align-items: center;">
						<h1>{{APP_NAME}}</h1>
						<i class="iconfont success">&#xe617;</i>
					</div>
					<p style="padding: 20px 0;">{{APP_NAME}} has been installed successfully to your workspace, Thank you!</p>
				</div>
			</div>
		</body>
		</html>
	`
)
