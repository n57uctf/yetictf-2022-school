from flask import Flask, render_template, make_response, redirect
import random

app = Flask(__name__)



@app.route('/')
def index():
	flag1 = random.choice(['/YWxwaGE','/YmV0YQ','/Z2FtbWE','/ZGVsdGE'])
	flag2 = random.choice(['/YWxwaGE','/YmV0YQ','/Z2FtbWE','/ZGVsdGE'])
	flag3 = random.choice(['/YWxwaGE','/YmV0YQ','/Z2FtbWE','/ZGVsdGE'])
	flag4 = random.choice(['/YWxwaGE','/YmV0YQ','/Z2FtbWE','/ZGVsdGE'])
	style1 = random.choice(['btn btn-lg btn-success','btn btn-lg btn-info','btn btn-lg btn-warning','btn btn-lg btn-danger'])
	style2 = random.choice(['btn btn-lg btn-success','btn btn-lg btn-info','btn btn-lg btn-warning','btn btn-lg btn-danger'])
	style3 = random.choice(['btn btn-lg btn-success','btn btn-lg btn-info','btn btn-lg btn-warning','btn btn-lg btn-danger'])
	style4 = random.choice(['btn btn-lg btn-success','btn btn-lg btn-info','btn btn-lg btn-warning','btn btn-lg btn-danger'])
	resp = make_response(render_template('index.html', flag1=flag1,flag2=flag2,flag3=flag3,flag4=flag4,style1=style1,style2=style2,style3=style3,style4=style4))
	resp.set_cookie('SESSION_IDI', '', expires=0)
	return resp

@app.route('/YWxwaGE')
def flag1():
	response = make_response(redirect('/'))
	response.set_cookie('SESSION_IDI','3x3_')
	return response

@app.route('/YmV0YQ')
def flag2():
	response = make_response(redirect('/'))
	response.set_cookie('SESSION_IDI','cu73_')
	return response

@app.route('/Z2FtbWE')
def flag3():
	response = make_response(redirect('/'))
	response.set_cookie('SESSION_IDI','Cr3a73_')
	return response

@app.route('/ZGVsdGE')
def flag4():
	response = make_response(redirect('/'))
	response.set_cookie('SESSION_IDI','7H1S_R3V4_@')
	return response



if __name__ == '__main__':
		app.run(host='0.0.0.0',port=8087)
