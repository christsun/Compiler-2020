#!python3
from flask_wtf import FlaskForm
from wtforms.validators import ValidationError, DataRequired, Email, EqualTo
from wtforms import *


class LoginForm(FlaskForm):
    userid = StringField(
        'StuID', validators=[DataRequired()], render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever'
        }
    )

    userpassword = PasswordField(
        'Password', validators=[DataRequired()], render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever'
        }
    )

    submit = SubmitField(
        'Login', render_kw={
            'class': "btn btn-primary"
        }
    )


class RegistrationForm(FlaskForm):
    userid = StringField(
        'StuID', validators=[DataRequired()],
        render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )
    password = PasswordField(
        'Password',
        validators=[DataRequired()],
        render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )
    password2 = PasswordField(
        'Repeat Password',
        validators=[DataRequired(), EqualTo('password')],
        render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )
    email = StringField(
        'Email',
        validators=[DataRequired(), Email()],
        render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )
    student_name = StringField(
        'Your real name',
        validators=[DataRequired()],
        render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )
    repo_url = StringField(
        'Repo',
        validators=[DataRequired()],
        render_kw = {
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )
    submit = SubmitField(
        'Register',
        render_kw={
            'class': "btn btn-primary",
        }
    )


class ModifyPasswordForm(FlaskForm):
    old_password = PasswordField(
        'Old Password',
        validators=[DataRequired()],
        render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )

    password = PasswordField(
        'New Password',
        validators=[DataRequired()],
        render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )

    password2 = PasswordField(
        'Repeat Password',
        validators=[DataRequired(), EqualTo('password')],
        render_kw={
            'class': 'form-control monospace',
            'placeholder': 'Whatever',
        }
    )

    submit = SubmitField(
        'Confirm',
        render_kw={
            'class': "btn btn-primary",
        }
    )
