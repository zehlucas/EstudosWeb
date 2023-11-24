from flask import Flask, render_template

app = Flask(__name__)

@app.route('/inicio')
def ola():
    return render_template('lista.html')

app.run(
    host='0.0.0.0',
    port=80
)